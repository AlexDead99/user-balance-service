// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: product.sql

package db

import (
	"context"
)

const getProduct = `-- name: GetProduct :one
SELECT product_id, name, price, amount FROM products
WHERE product_id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, productID int32) (Products, error) {
	row := q.db.QueryRowContext(ctx, getProduct, productID)
	var i Products
	err := row.Scan(
		&i.ProductID,
		&i.Name,
		&i.Price,
		&i.Amount,
	)
	return i, err
}

const getProductForUpdate = `-- name: GetProductForUpdate :one
SELECT product_id, name, price, amount FROM products
WHERE product_id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetProductForUpdate(ctx context.Context, productID int32) (Products, error) {
	row := q.db.QueryRowContext(ctx, getProductForUpdate, productID)
	var i Products
	err := row.Scan(
		&i.ProductID,
		&i.Name,
		&i.Price,
		&i.Amount,
	)
	return i, err
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
  set amount = $2
WHERE product_id = $1
RETURNING product_id, name, price, amount
`

type UpdateProductParams struct {
	ProductID int32 `json:"product_id"`
	Amount    int32 `json:"amount"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Products, error) {
	row := q.db.QueryRowContext(ctx, updateProduct, arg.ProductID, arg.Amount)
	var i Products
	err := row.Scan(
		&i.ProductID,
		&i.Name,
		&i.Price,
		&i.Amount,
	)
	return i, err
}
