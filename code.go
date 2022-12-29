package main

import (
	"database/sql"
	"fmt"
	"log"

	//"context"
	//"time"
	//"strconv"
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


    err = CreateProductTable(db)
    if err != nil {
        log.Printf("Create product table failed with error %s", err)
        return
    }
	router.GET("/products", GetProducts)
	//  router.GET("/products/:id", getProductsByCatg)
	router.GET("/product-:id", GetProductByID)
	router.POST("/addproduct", PostProducts)
	router.GET("/deleteproduct-:id", DeleteProductByID)


	err = CreateInventoryTable(db)
    if err != nil {
        log.Printf("Create product table failed with error %s", err)
        return
    }
    router.GET("/inventorys", GetInventorys)
	router.GET("/inventory-:id", GetInventoryByID)
	router.POST("/addinventory", PostInventorys)
	router.GET("/deleteinventory-:id", DeleteInventoryByID)


	err = CreateCategoryTable(db)
    if err != nil {
        log.Printf("Create product table failed with error %s", err)
        return
    }
    router.GET("/categorys", GetCategorys)
	router.GET("/category-:id", GetCategoryByID)
	router.POST("/addcategory", PostCategorys)
	router.GET("/deletecategory-:id", DeleteCategoryByID)


	err = CreateCartTable(db)
    if err != nil {
        log.Printf("Create product table failed with error %s", err)
        return
    }
	router.GET("/carts", GetCarts)
	router.GET("/cart-:id", GetCartByID)
	router.POST("/addcart", PostCarts)
	router.GET("/deletecart-:id", DeleteCartByID)


	router.Run("localhost:8080")
}
