package app

import (
	"database/sql"
	"fmt"
	"minang-kos-service/helper"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	godotenv.Load()
	driverName := os.Getenv("DB_DRIVER")

	db, err := sql.Open(driverName, getDataSourceName())
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func getDataSourceName() string {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", dbUsername, dbPassword, dbHost, dbName)
}
