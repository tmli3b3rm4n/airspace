package database

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Database interface {
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	Preload(query string, args ...interface{}) (tx *gorm.DB)
	Model(value interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Limit(limit int) (tx *gorm.DB)
	Order(value interface{}) (tx *gorm.DB)
	Select(query interface{}, args ...interface{}) (tx *gorm.DB)
	Save(value interface{}) (tx *gorm.DB)
	Row() *sql.Row
	Scan(dest interface{}) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error)
	Table(name string, args ...interface{}) (tx *gorm.DB)
	AutoMigrate(dst ...interface{}) error
	Begin(opts ...*sql.TxOptions) *gorm.DB
	Commit() *gorm.DB
	Update(column string, value interface{}) (tx *gorm.DB)
	Rollback() *gorm.DB
	Exec(sql string, values ...interface{}) *gorm.DB
	WithContext(ctx context.Context) (tx *gorm.DB)
}

const LOCAL = "local"

func Connect() (Database, error) {
	// Get environment variables for user and password
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	dbname := os.Getenv("POSTGRES_DB")

	if user == "" || password == "" || host == "" || dbname == "" {
		log.Fatal("Missing necessary environment variables")
	}

	// Construct the Data Source Name (DSN)
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=5432 sslmode=disable", user, password, dbname, host)

	// Connect to the PostgresSQL database using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
