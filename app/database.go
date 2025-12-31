package app

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/kevinjuliow/dataInventarisBarang/helper"
)

func LoadEnv() {
	err := godotenv.Load()
	helper.PanicIfError(err)
}

func NewDb() *sql.DB {
	LoadEnv()
	dbName := os.Getenv("DB_USERNAME")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbUrl := dbName + ":" + dbpassword + "@tcp(localhost:3306)/data_inventaris"
	db, err := sql.Open("mysql", dbUrl)
	helper.PanicIfError(err)

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
