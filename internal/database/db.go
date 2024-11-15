package database

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
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
	Update(model interface{}) error
	Rollback() *gorm.DB
	Exec(sql string, values ...interface{}) *gorm.DB
}

const LOCAL = "local"

func Connect() (Database, error) {
	var db *gorm.DB
	var err error

	if os.Getenv("ENV") == LOCAL {
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	} else {
		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"), "5432", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PASSWORD"))
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Unable to access gorm db. Error: %v", err)
		return nil, err
	} else {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	return db, nil
}
