package helper

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Init creates a connection to mysql database and
// migrates any new models
func Init() {
	user := getEnv("PG_USER", "hugo")
	password := getEnv("PG_PASSWORD", "")
	host := getEnv("PG_HOST", "localhost")
	port := getEnv("PG_PORT", "8080")
	database := getEnv("PG_DB", "products")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

//CloseDB ...
func CloseDB() {
	db.Close()
}
