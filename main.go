package main

import (
	"database/sql"
	"log"

	"github.com/AlexDead99/user-balance-service/api"
	db "github.com/AlexDead99/user-balance-service/db/sqlc"
	"github.com/AlexDead99/user-balance-service/docs"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://test:test@postgres:5432/simple?sslmode=disable"
	serverAddress = "0.0.0.0:3000"
)

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

}

// @title           Swagger Balance API
// @version         1.0
// @description     This is a small user balance server.
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
func main() {

	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot establish connection to database: ", err)
	}
	runDBMigration("file://db/migration", dbSource)

	err = conn.Ping()
	if err != nil {
		log.Fatal("cannot establish connection to database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start: ", err)
	}

}
