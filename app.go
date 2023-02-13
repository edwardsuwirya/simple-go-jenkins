package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	run(port, dsn)
}

func initRouter(route *gin.Engine, db *sqlx.DB) {
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.GET("/products", func(c *gin.Context) {
		var products []Product
		err := db.Select(&products, "SELECT p.id,p.product_name,p.category_id,pc.category_name FROM m_product p join m_product_category pc on p.category_id = pc.id")
		if err != nil {
			fmt.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"message": products,
		})
	})
}

func run(port string, dsn string) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	route := gin.Default()
	initRouter(route, db)
	err = route.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Server is not running %s", err)
	}
}
