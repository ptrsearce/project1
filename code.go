package main

import (
	"database/sql"
	//"strconv"
	"fmt"
	"log"

	// "os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	router := gin.Default()

	// Capture connection properties.
	cfg := mysql.Config{
		// User:   os.Getenv("DBUSER"),
		// Passwd: os.Getenv("DBPASS"),
		User:   "root",
		Passwd: "Teja@1102",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "dbproj",
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Database Connected!")

	// pdtID, err := addProduct(product{
	//     Product_id: 100
	//     Name: "cello"
	//     Spec: {"colour":"blue","height_in_cm":90}
	//     Catg_id: "5"
	//     Price  150.0
	// })

	router.GET("/products", GetProducts)
	//  router.GET("/products/:id", getProductsByCatg)
	router.GET("/products/:id", GetProductByID)
	router.POST("/products", PostProducts)

    router.GET("/inventorys", GetInventorys)
	router.GET("/inventorys/:id", GetInventoryByID)
	router.POST("/inventorys", PostInventorys)

    router.GET("/categorys", GetCategorys)
	router.GET("/categorys/:id", GetCategoryByID)
	router.POST("/categorys", PostCategorys)

	router.Run("localhost:8080")
}
