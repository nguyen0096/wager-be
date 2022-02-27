package cmd

// import (
// 	"database/sql"
// 	"fmt"

// 	"github.com/golang-migrate/migrate/v4"
// 	"github.com/golang-migrate/migrate/v4/database/postgres"
// 	_ "github.com/golang-migrate/migrate/v4/source/file"
// 	_ "github.com/lib/pq"
// )

// func initSchema() {
// 	db, err := sql.Open("postgres", "postgres://wager:wager@localhost:5444/wagerdb?sslmode=disable")
// 	if err != nil {
// 		panic(fmt.Errorf("migrate failed. err: %w", err))
// 	}
// 	driver, err := postgres.WithInstance(db, &postgres.Config{})
// 	if err != nil {
// 		panic(fmt.Errorf("migrate failed. err: %w", err))
// 	}

// 	m, err := migrate.NewWithDatabaseInstance(
// 		"file://db/migrations",
// 		"postgres", driver)
// 	if err != nil {
// 		panic(fmt.Errorf("migrate failed. err: %w", err))
// 	}

// 	m.Up()
// }
