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
