package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"os"
)

type StorageConfig struct {
	username, password, host, port, database string
}

func init() {
	err := godotenv.Load("./repository/.env")
	if err != nil {
		fmt.Println("unable to download .env")
	}

}

func NewClient() (*pgx.Conn, error) {
	var sc = StorageConfig{
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("DATABASE"),
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", sc.host, sc.port, sc.username, sc.password, sc.database)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	//defer conn.Close(context.Background())

	fmt.Println("DATABASE CONNECT")
	return conn, err
}
