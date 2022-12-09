package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type StorageConfig struct {
	username, password, host, port, database string
}

func NewClient() {
	var sc = StorageConfig{
		"admin",
		"root",
		"localhost",
		"5432",
		"db_cinema",
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", sc.host, sc.port, sc.username, sc.password, sc.database)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		return
	}
	defer conn.Close(context.Background())

	fmt.Println("DATABASE CONNECT")
}
