package models

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func SetDB(gdb *gorm.DB) {
	db = gdb
}

func InitDB() {
	e := godotenv.Load(".env")
	if e != nil {
		log.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbMaxIdleConns, _ := strconv.Atoi(os.Getenv("db_max_idle_conns"))
	dbMaxOpenConns, _ := strconv.Atoi(os.Getenv("db_max_open_conns"))

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Print(err)
	}

	db = conn
	db.DB().SetMaxIdleConns(dbMaxIdleConns)
	db.DB().SetMaxOpenConns(dbMaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Hour)
}

func GetDB() *gorm.DB {
	return db
}
