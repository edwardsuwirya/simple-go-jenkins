package main

import (
	"enigmacamp.com/godocker/dto"
	"enigmacamp.com/godocker/models"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	migration := flag.Bool("migration", false, "DB Migration")
	flag.Parse()
	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Database connected")

	if *migration {
		err := db.AutoMigrate(&models.Product{}, &models.ProductCategory{})
		if err != nil {
			return
		}
		cat01 := models.ProductCategory{ID: "008", CategoryName: "Appetizer"}
		cat02 := models.ProductCategory{ID: "009", CategoryName: "Dessert"}
		categories := []models.ProductCategory{cat01, cat02}
		db.Create(categories)
		prod01 := models.Product{ID: "003", ProductName: "Cream Soup", ProductCategory: cat01}
		prod02 := models.Product{ID: "004", ProductName: "1 scoop ice cream", ProductCategory: cat02}
		prods := []models.Product{prod01, prod02}
		db.Create(prods)
	} else {
		route := gin.Default()
		route.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		route.GET("/products", func(c *gin.Context) {
			var products []dto.Product
			db.Model(&models.ProductCategory{}).Preload("Products").Find(&products)
			err := db.Model(&models.Product{}).Select("m_product.id", "m_product.product_name", "m_product.category_id", "m_product_category.category_name").Joins("JOIN m_product_category on m_product_category.id=m_product.category_id").Find(&products).Error
			if err != nil {
				fmt.Println(err)
				return
			}
			c.JSON(200, gin.H{
				"message": products,
			})
		})

		err = route.Run(fmt.Sprintf(":%s", port))
		if err != nil {
			log.Fatalf("Server is not running %s", err)
		}
	}

}
