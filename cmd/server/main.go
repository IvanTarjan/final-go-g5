package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/IvanTarjan/final-go-g5/cmd/server/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := connectDB()

	engine := gin.New()
	engine.Use(gin.Recovery())

	runApp(db, engine)

	defer db.Close()
}

func runApp(db *sql.DB, engine *gin.Engine) {
	router := router.NewRouter(engine, db)
	router.MapRoutes()
	if err := engine.Run("localhost:8080"); err != nil {
		panic(err)
	}
}

func connectDB() *sql.DB {
	var dbUser, dbPass, dbHost, dbPort, dbName string
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
