package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("error in transaction: %v rollback error: %v", err, rollbackErr)
		}
		return err
	}

	return tx.Commit()
}

type ProductsParams struct {
	ProductId int32 `json:"product_id"`
	Amount    int32 `json:"amount"`
}
type TransferTxParams struct {
	UserId      int32             `json:"user_id"`
	Products    []*ProductsParams `json:"products"`
	ServiceId   int32             `json:"service_id"`
	Description string            `json:"description"`
}

type TransferTxResult struct {
	Success bool  `json:"success"`
	Id      int32 `json:"transaction_id"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var totalPrice float32 = 0.0

		for _, product := range arg.Products {
			productInfo, err := q.GetProductForUpdate(context.Background(), product.ProductId)
			if err != nil {
				return err
			}

			availableAmount := productInfo.Amount - product.Amount
			if availableAmount < 0 {
				return fmt.Errorf("You can not reserved so much products")
			}

			updateParams := UpdateProductParams{
				ProductID: product.ProductId,
				Amount:    availableAmount,
			}
			_, err = q.UpdateProduct(context.Background(), updateParams)
			if err != nil {
				return err
			}

			totalPrice += productInfo.Price * float32(product.Amount)
		}

		user, err := q.GetAccount(context.Background(), arg.UserId)
		if err != nil {
			return err
		}

		availableMoney := user.Balance - totalPrice
		if availableMoney < 0 {
			return fmt.Errorf("You can't pay for this order. Check your balance")
		}

		updateUser := UpdateAccountParams{
			AccountID: arg.UserId,
			Balance:   availableMoney,
		}
		_, err = q.UpdateAccount(context.Background(), updateUser)
		if err != nil {
			return err
		}

		createTransfer := CreateTransferParams{
			UserID:      arg.UserId,
			ServiceID:   arg.ServiceId,
			TotalPrice:  totalPrice,
			Description: arg.Description,
			Status:      "Pending",
		}
		transfer, err := q.CreateTransfer(context.Background(), createTransfer)

		if err != nil {
			return err
		}
		for _, product := range arg.Products {
			orderParams := CreateOrderParams{
				TransferID: transfer.TransferID,
				ProductID:  product.ProductId,
				Amount:     product.Amount,
			}
			_, err := q.CreateOrder(context.Background(), orderParams)

			if err != nil {
				return err
			}
		}

		result.Success = true
		result.Id = transfer.TransferID

		return nil

	})

	return result, err
}

type UpdateUserBalanceTxParams struct {
	UserId int32   `json:"user_id"`
	Amount float32 `json:"amount"`
}
type UpdateUserBalanceTxResult struct {
	UserId  int32 `json:"user_id"`
	Success bool  `json:"success"`
}

func (store *Store) UpdateUserBalanceTx(ctx context.Context, arg UpdateUserBalanceTxParams) (UpdateUserBalanceTxResult, error) {
	var result UpdateUserBalanceTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		user, err := q.GetAccountForUpdate(ctx, arg.UserId)
		if err != nil {
			return err
		}
		availableMoney := user.Balance + arg.Amount

		if availableMoney < 0 {
			return fmt.Errorf("Balance can't be negative")
		}

		serviceId := 1 //Buy
		description := "Depositing funds to the wallet"

		if arg.Amount < 0 {
			serviceId = 2 //Sell
			description = "Withdrawal of funds from the wallet"

		}

		transferParams := CreateTransferParams{
			UserID:      arg.UserId,
			ServiceID:   int32(serviceId),
			TotalPrice:  availableMoney,
			Description: description,
			Status:      "Success",
		}
		transfer, err := q.CreateTransfer(ctx, transferParams)

		if err != nil {
			return err
		}

		userParams := UpdateAccountParams{
			AccountID: arg.UserId,
			Balance:   transfer.TotalPrice,
		}
		updatedUser, err := q.UpdateAccount(ctx, userParams)
		if err != nil {
			return err
		}

		result.Success = true
		result.UserId = updatedUser.AccountID

		return nil
	})

	return result, err
}

type DeleteTransferTxParams struct {
	TransferId int32 `json:"transfer_id"`
}
type DeleteTransferTxResult struct {
	Success bool `json:"success"`
}

func (store *Store) DeleteTransferTx(ctx context.Context, arg DeleteTransferTxParams) (DeleteTransferTxResult, error) {
	var result DeleteTransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		transfer, err := q.GetTransfer(ctx, arg.TransferId)
		if err != nil {
			return err
		}

		orders, err := q.ListOrders(ctx, transfer.TransferID)
		if err != nil {
			return err
		}

		products := []*ProductsParams{}

		for _, order := range orders {
			product := &ProductsParams{
				ProductId: order.ProductID,
				Amount:    order.Amount,
			}
			products = append(products, product)
		}

		for _, product := range products {
			productForUpdate, err := q.GetProductForUpdate(ctx, product.ProductId)
			if err != nil {
				return err
			}

			updateParam := UpdateProductParams{
				ProductID: product.ProductId,
				Amount:    productForUpdate.Amount + product.Amount,
			}

			_, err = q.UpdateProduct(ctx, updateParam)
			if err != nil {
				return err
			}
		}

		user, err := q.GetAccountForUpdate(ctx, transfer.UserID)
		if err != nil {
			return err
		}

		userParams := UpdateAccountParams{
			AccountID: user.AccountID,
			Balance:   user.Balance + transfer.TotalPrice,
		}

		_, err = q.UpdateAccount(ctx, userParams)
		if err != nil {
			return err
		}

		for _, order := range orders {
			q.DeleteOrder(ctx, order.OrderID)
		}

		updateTransferParam := UpdateTransferParams{
			TransferID: transfer.TransferID,
			Status:     "Failed",
		}
		q.UpdateTransfer(ctx, updateTransferParam)

		result.Success = true
		return nil

	})

	return result, err

}
