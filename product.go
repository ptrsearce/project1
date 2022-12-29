package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
)

type product struct {
	Product_id int    `json:"id"`
	Name       string `json:"title"`
	Spec       string `json:"artist"`
	Catg_id    int
	Price      float64 `json:"price"`
}

func CreateProductTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS product (
		product_id int unsigned NOT NULL,
		name varchar(45) DEFAULT NULL,
		spec json DEFAULT NULL,
		catg_id int DEFAULT NULL,
		price float DEFAULT NULL,
		PRIMARY KEY (product_id),
		UNIQUE KEY product_id_UNIQUE (product_id) /*!80000 INVISIBLE */,
		KEY catg_id_idx (catg_id),
		CONSTRAINT catg_id FOREIGN KEY (catg_id) REFERENCES category (category_id)
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

func GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var pdt product

	row := db.QueryRow("SELECT * FROM product WHERE product_id = ?", id)
	if err := row.Scan(&pdt.Product_id, &pdt.Name, &pdt.Spec, &pdt.Catg_id, &pdt.Price); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
			fmt.Errorf("productsById %d: no such product", id)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
		fmt.Errorf("productsById %d: %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, pdt)

}

// getProducts responds with the list of all products as JSON.
func GetProducts(c *gin.Context) {

	var products []product

	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		fmt.Errorf("products : %v", err)
		return
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var pdt product
		if err := rows.Scan(&pdt.Product_id, &pdt.Name, &pdt.Spec, &pdt.Catg_id, &pdt.Price); err != nil {
			fmt.Errorf("products : %v", err)
			return
		}
		products = append(products, pdt)
	}
	if err := rows.Err(); err != nil {
		fmt.Errorf("products : %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, products)
}

// getProductByID locates the product whose ID value matches the id
// parameter sent by the client, then returns that product as a response.

// postProducts adds an product from JSON received in the request body.
func PostProducts(c *gin.Context) {
	var newProduct product

	// Call BindJSON to bind the received JSON to
	// newProduct.
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	result, err := db.Exec("INSERT INTO product (product_id, name, spec, catg_id, price) VALUES (?, ?, ?, ?, ?)", newProduct.Product_id, newProduct.Name, newProduct.Spec, newProduct.Catg_id, newProduct.Price)
	// Add the new product to the slice.
	if err != nil {
		fmt.Errorf("addProduct: %v", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Errorf("addProduct: %v", err)
		return
	}
	fmt.Print(id)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func DeleteProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var pdt product

	row := db.QueryRow("DELETE FROM product WHERE product_id = ?", id)
	if err := row.Scan(&pdt.Product_id, &pdt.Name, &pdt.Spec, &pdt.Catg_id, &pdt.Price); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
			fmt.Errorf("productsById %d: no such product", id)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
		fmt.Errorf("productsById %d: %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, pdt)

}



// &pdt.Product_id, &pdt.Name, &pdt.Spec, &pdt.Catg_id, &pdt.Price

// DELETE FROM `table_name` [WHERE condition];
