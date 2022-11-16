// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: transfer.sql

package db

import (
	"context"
	"time"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
   user_id, service_id, total_price, description, status
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING transfer_id, user_id, service_id, total_price, created_at, description, status
`

type CreateTransferParams struct {
	UserID      int32   `json:"user_id"`
	ServiceID   int32   `json:"service_id"`
	TotalPrice  float32 `json:"total_price"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfers, error) {
	row := q.db.QueryRowContext(ctx, createTransfer,
		arg.UserID,
		arg.ServiceID,
		arg.TotalPrice,
		arg.Description,
		arg.Status,
	)
	var i Transfers
	err := row.Scan(
		&i.TransferID,
		&i.UserID,
		&i.ServiceID,
		&i.TotalPrice,
		&i.CreatedAt,
		&i.Description,
		&i.Status,
	)
	return i, err
}

const generalReport = `-- name: GeneralReport :many
SELECT transfer_id, user_id, transfers.service_id, total_price, created_at, description, status, services.service_id, name FROM transfers
JOIN services
ON transfers.service_id = services.service_id
WHERE EXTRACT(MONTH FROM created_at) = $1 AND EXTRACT(YEAR FROM created_at) = $2 AND status IN ('Success')
`

type GeneralReportParams struct {
	CreatedAt   time.Time `json:"created_at"`
	CreatedAt_2 time.Time `json:"created_at_2"`
}

type GeneralReportRow struct {
	TransferID  int32     `json:"transfer_id"`
	UserID      int32     `json:"user_id"`
	ServiceID   int32     `json:"service_id"`
	TotalPrice  float32   `json:"total_price"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	ServiceID_2 int32     `json:"service_id_2"`
	Name        string    `json:"name"`
}

func (q *Queries) GeneralReport(ctx context.Context, arg GeneralReportParams) ([]GeneralReportRow, error) {
	rows, err := q.db.QueryContext(ctx, generalReport, arg.CreatedAt, arg.CreatedAt_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GeneralReportRow
	for rows.Next() {
		var i GeneralReportRow
		if err := rows.Scan(
			&i.TransferID,
			&i.UserID,
			&i.ServiceID,
			&i.TotalPrice,
			&i.CreatedAt,
			&i.Description,
			&i.Status,
			&i.ServiceID_2,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTransfer = `-- name: GetTransfer :one
SELECT transfer_id, user_id, service_id, total_price, created_at, description, status FROM transfers
WHERE transfer_id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, transferID int32) (Transfers, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, transferID)
	var i Transfers
	err := row.Scan(
		&i.TransferID,
		&i.UserID,
		&i.ServiceID,
		&i.TotalPrice,
		&i.CreatedAt,
		&i.Description,
		&i.Status,
	)
	return i, err
}

const transactionHistoryForUser = `-- name: TransactionHistoryForUser :many
SELECT transfer_id, user_id, service_id, total_price, created_at, description, status FROM transfers
WHERE user_id = $1 AND EXTRACT(MONTH FROM created_at) >= $1 AND EXTRACT(YEAR FROM created_at) >= $2 AND status IN ('Success')
`

type TransactionHistoryForUserParams struct {
	UserID    int32     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) TransactionHistoryForUser(ctx context.Context, arg TransactionHistoryForUserParams) ([]Transfers, error) {
	rows, err := q.db.QueryContext(ctx, transactionHistoryForUser, arg.UserID, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfers
	for rows.Next() {
		var i Transfers
		if err := rows.Scan(
			&i.TransferID,
			&i.UserID,
			&i.ServiceID,
			&i.TotalPrice,
			&i.CreatedAt,
			&i.Description,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransfer = `-- name: UpdateTransfer :one
UPDATE transfers
  set status = $2
WHERE transfer_id = $1
RETURNING transfer_id, user_id, service_id, total_price, created_at, description, status
`

type UpdateTransferParams struct {
	TransferID int32  `json:"transfer_id"`
	Status     string `json:"status"`
}

func (q *Queries) UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (Transfers, error) {
	row := q.db.QueryRowContext(ctx, updateTransfer, arg.TransferID, arg.Status)
	var i Transfers
	err := row.Scan(
		&i.TransferID,
		&i.UserID,
		&i.ServiceID,
		&i.TotalPrice,
		&i.CreatedAt,
		&i.Description,
		&i.Status,
	)
	return i, err
}
