#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/db/migration -database postgresql://test:test@postgres:5432/simple?sslmode=disable -verbose up
