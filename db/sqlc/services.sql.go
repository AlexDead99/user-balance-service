// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: services.sql

package db

import (
	"context"
)

const getService = `-- name: GetService :one
SELECT service_id, name FROM services
WHERE service_id = $1 LIMIT 1
`

func (q *Queries) GetService(ctx context.Context, serviceID int32) (Services, error) {
	row := q.db.QueryRowContext(ctx, getService, serviceID)
	var i Services
	err := row.Scan(&i.ServiceID, &i.Name)
	return i, err
}
