package db

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
    connStr := "user=junin password=jzin007 dbname=shopping_db sslmode=disable"

    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        return fmt.Errorf("error connecting to database: %v", err)
    }

    if err = DB.Ping(); err != nil {
        return fmt.Errorf("cannot ping database: %v", err)
    }

    fmt.Println("Connected to database!")
    return nil
}