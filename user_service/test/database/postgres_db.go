package test_database

import (
	"database/sql"
	"fmt"
	"log"
	"user_service/src/common/database"
)

type PostgresDb struct {
	db *sql.DB
}

func NewPostgresDb() *PostgresDb {
	db, _ := database.NewPostgresUserDatabase(database.DbConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "admin",
		Password: "admin123",
		DBName:   "user_db",
		SSLMode:  "disable",
	})

	return &PostgresDb{
		db: db,
	}
}

func (r *PostgresDb) GetDb() *sql.DB {
	return r.db
}

func (r *PostgresDb) ClearTable(tableName string) {
	query := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tableName)

	_, err := r.db.Exec(query)
	if err != nil {
		log.Print(fmt.Errorf("failed to clear table %s: %w", tableName, err))
	}
}

func (r *PostgresDb) GetCount(tableName string) int {
	var rowCount int
	err := r.db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)).Scan(&rowCount)
	if err != nil {
		log.Fatal(err)
	}

	return rowCount
}

func (r *PostgresDb) GetAll(tableName string) *sql.Rows {
	rows, err := r.db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		log.Fatal(err)
	}

	return rows
}
