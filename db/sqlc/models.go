// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"
)

type Accounts struct {
	AccountID int32          `json:"account_id"`
	Owner     sql.NullString `json:"owner"`
	Balance   sql.NullString `json:"balance"`
	CreatedAt time.Time      `json:"created_at"`
}

type OrdersDetails struct {
	OrderID    int32         `json:"order_id"`
	TransferID sql.NullInt32 `json:"transfer_id"`
	ProductID  sql.NullInt32 `json:"product_id"`
	Amount     sql.NullInt32 `json:"amount"`
	CreatedAt  time.Time     `json:"created_at"`
}

type Products struct {
	ProductID int32          `json:"product_id"`
	Name      sql.NullString `json:"name"`
	Price     sql.NullString `json:"price"`
	Amount    sql.NullInt32  `json:"amount"`
}

type Services struct {
	ServiceID int32          `json:"service_id"`
	Name      sql.NullString `json:"name"`
}

type Transfers struct {
	TransferID int32          `json:"transfer_id"`
	UserID     sql.NullInt32  `json:"user_id"`
	ServiceID  sql.NullInt32  `json:"service_id"`
	TotalPrice sql.NullString `json:"total_price"`
	CreatedAt  time.Time      `json:"created_at"`
}
