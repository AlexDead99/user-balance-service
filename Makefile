POSTGRES_USER=postgres
POSTGRES_PASS=root
POSTGRES_URL=localhost:5432
POSTGRES_DB=simple

up:
	migrate -path db/migration -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASS)@$(POSTGRES_URL)/$(POSTGRES_DB)?sslmode=disable" -verbose up
down:
	migrate -path db/migration -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASS)@$(POSTGRES_URL)/$(POSTGRES_DB)?sslmode=disable" -verbose down
