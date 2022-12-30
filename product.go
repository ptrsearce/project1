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
	Product_id int    `json:"product_id"`
	Name       string `json:"name"`
	Spec       string `json:"spec"`
	Catg_id    int	  `json:"catg_id"`
	Price      float64 `json:"price"`
}

type productsimple struct {
	Product_id int    `json:"product_id"`
	Name       string `json:"name"`
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

	var products []productsimple

	rows, err := db.Query("SELECT product_id,name,price FROM product")
	if err != nil {
		fmt.Errorf("products : %v", err)
		return
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var pdt productsimple
		if err := rows.Scan(&pdt.Product_id, &pdt.Name, &pdt.Price); err != nil {
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
	db.QueryRow("DELETE FROM product WHERE product_id = ?", id)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "product is deleted"})

}


func UpdateProductByID(c *gin.Context) {
	var newProduct product
	id, _ := strconv.Atoi(c.Param("id"))
	// Call BindJSON to bind the received JSON to
	// newProduct.
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	if(newProduct.Product_id !=0){db.Exec("UPDATE product SET product_id= ? WHERE product_id = ?",newProduct.Product_id, id)}
	if(newProduct.Name !=""){db.Exec("UPDATE product SET name= ? WHERE product_id = ?",newProduct.Name, id)}
	if(newProduct.Spec !=""){db.Exec("UPDATE product SET spec= ? WHERE product_id = ?",newProduct.Spec, id)}
	if(newProduct.Catg_id !=0){db.Exec("UPDATE product SET catg_id= ? WHERE product_id = ?",newProduct.Catg_id, id)}
	if(newProduct.Price !=0){db.Exec("UPDATE product SET price= ? WHERE product_id = ?",newProduct.Price, id)}
	//result, err := db.Exec("UPDATE product  SET (product_id, name, spec, catg_id, price)= (newProduct.Product_id, newProduct.Name, newProduct.Spec, newProduct.Catg_id, newProduct.Price) WHERE product_id = ?",newProduct.Product_id, newProduct.Name, newProduct.Spec, newProduct.Catg_id, newProduct.Price, id)
	// Add the new product to the slice.
	// if err != nil {
	// 	fmt.Errorf("addProduct: %v", err)
	// 	return
	// }
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	fmt.Errorf("addProduct: %v", err)
	// 	return
	// }
	fmt.Print(id)
	c.IndentedJSON(http.StatusCreated, newProduct)

}

// &pdt.Product_id, &pdt.Name, &pdt.Spec, &pdt.Catg_id, &pdt.Price

// DELETE FROM `table_name` [WHERE condition];
