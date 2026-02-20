package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func SqlxInitDB(dbConfig DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", getConnectionStr(dbConfig))
	if err != nil {
		return nil, fmt.Errorf("could not connect to the database, err:\n %+v", err)
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping the database, err:\n %+v", err)
	}
	return db, nil
}

func PgxInitDB(ctx context.Context, dbConfig DBConfig) (*pgx.Conn, error) {
	db, err := pgx.Connect(ctx, getConnectionStr(dbConfig))
	if err != nil {
		return nil, fmt.Errorf("could not connect to the database, err:\n %+v", err)
	}
	return db, nil
}

func getConnectionStr(dbConfig DBConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Name,
	)
}
