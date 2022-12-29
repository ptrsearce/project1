package main

import (
	"strconv"
	"net/http"
	"database/sql"
	"context"
	"time"
	"log"
    
    "fmt"

    "github.com/gin-gonic/gin"
)

type cart struct {
	Cart_id int `json:"id"`
	Product_id int
	Quantity     int
	Price float32
}

func CreateCartTable(db *sql.DB) error {  
    query :=`CREATE TABLE IF NOT EXISTS cart (
		card_id int NOT NULL,
		product_id varchar(45) DEFAULT NULL,
		quantity int DEFAULT NULL,
		price float DEFAULT NULL,
		PRIMARY KEY (card_id)
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci`

    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    res, err := db.ExecContext(ctx, query)
    if err != nil {
        log.Printf("Error %s when creating product table", err)
        return err
    }
    rows, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error %s when getting rows affected", err)
        return err
    }
    log.Printf("Rows affected when creating table: %d", rows)
    return nil
}


func GetCartByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var crt cart

	row := db.QueryRow("SELECT * FROM cart WHERE cart_id = ?", id)
	if err := row.Scan(&crt.Cart_id, &crt.Product_id, &crt.Quantity, &crt.Price); err != nil {
		if err == sql.ErrNoRows {
		   c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cart not found"})
		   fmt.Errorf("cartsById %d: no such cart", id)
		   return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cart not found"})
		fmt.Errorf("cartsById %d: %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, crt)

}

// getCarts responds with the list of all carts as JSON.
func GetCarts(c *gin.Context) {

	var carts []cart

	rows, err := db.Query("SELECT * FROM cart")
	if err != nil {
		fmt.Errorf("carts : %v", err)
		return
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var crt cart
		if err := rows.Scan(&crt.Cart_id, &crt.Product_id, &crt.Quantity, &crt.Price); err != nil {
			fmt.Errorf("carts : %v", err)
			return
		}
		carts = append(carts, crt)
	}
	if err := rows.Err(); err != nil {
		fmt.Errorf("carts : %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, carts)
}

// getCartByID locates the cart whose ID value matches the id
// parameter sent by the client, then returns that cart as a response.

// postCarts adds an cart from JSON received in the request body.
func PostCarts(c *gin.Context) {
	var newCart cart

	// Call BindJSON to bind the received JSON to
	// newCart.
	if err := c.BindJSON(&newCart); err != nil {
		return
	}

	result, err := db.Exec("INSERT INTO cart (cart_id, product_id, quantity, price) VALUES (?, ?, ?, ?)", newCart.Cart_id, newCart.Product_id, newCart.Quantity, newCart.Price)
	// Add the new cart to the slice.
	if err != nil {
		fmt.Errorf("addCart: %v", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Errorf("addCart: %v", err)
		return
	}
	fmt.Print(id)
	c.IndentedJSON(http.StatusCreated, newCart)
}

func DeleteCartByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var crt cart

	row := db.QueryRow("DELETE FROM cart WHERE cart_id = ?", id)
	if err := row.Scan(&crt.Cart_id, &crt.Product_id, &crt.Quantity, &crt.Price); err != nil {
		if err == sql.ErrNoRows {
		   c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cart not found"})
		   fmt.Errorf("cartsById %d: no such cart", id)
		   return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cart not found"})
		fmt.Errorf("cartsById %d: %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, crt)

}