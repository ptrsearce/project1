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
	router.GET("/get_products", GetProducts)
	router.GET("/get_product_by_id-:id", GetProductByID)
	router.POST("/add_product", PostProducts)
	router.DELETE("/delete_product_by_id-:id", DeleteProductByID)
	router.PUT("/update_product_by_id-:id", UpdateProductByID)


	err = CreateInventoryTable(db)
    if err != nil {
        log.Printf("Create product table failed with error %s", err)
        return
    }
    router.GET("/get_inventory", GetInventory)
	router.GET("/get_inventory_by_id-:id", GetInventoryByID)
	router.POST("/add_inventory", PostInventory)
	router.DELETE("/delete_inventory_by_id-:id", DeleteInventoryByID)
	router.PUT("/update_inventory_by_id-:id", UpdateInventoryByID)


	err = CreateCategoryTable(db)
    if err != nil {
        log.Printf("Create product table failed with error %s", err)
        return
    }
    router.GET("/get_category", GetCategory)
	router.GET("/get_category_by_id-:id", GetCategoryByID)
	router.POST("/add_category", PostCategory)
	router.DELETE("/delete_category_by_id-:id", DeleteCategoryByID)
	router.PUT("/update_category_by_id-:id", UpdateCategoryByID)


	err = CreateCartTable(db)
    if err != nil {
        log.Printf("Create product table failed with error %s", err)
        return
    }
	router.GET("/get_cart", GetCart)
	router.GET("/get_cart_by_id-:id", GetCartByID)
	router.POST("/add_cart", PostCart)
	router.DELETE("/delete_cart_by_id-:id", DeleteCartByID)
	router.PUT("/update_cart_by_id-:id", UpdateCartByID)


	router.Run("localhost:8080")
}
