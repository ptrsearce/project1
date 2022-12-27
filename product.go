package main

import (
	"strconv"
	"net/http"
    
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


func GetProductByID(c *gin.Context) {
    id,_ := strconv.Atoi(c.Param("id"))

	var pdt product

    row := db.QueryRow("SELECT * FROM product WHERE product_id = ?", id)
    if err := row.Scan(&pdt.Product_id, &pdt.Name, &pdt.Spec, &pdt.Catg_id, &pdt.Price); err!=nil{
        // if err == sql.ErrNoRows {
        //    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
        //    fmt.Errorf("productsById %d: no such product", id) 
        //    return
        // }
        // c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
        // fmt.Errorf("productsById %d: %v", id, err)
        // return
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

// &pdt.Product_id, &pdt.Name, &pdt.Spec, &pdt.Catg_id, &pdt.Price