package main

import (
    "database/sql"
	"strconv"
    "fmt"
    "log"
    // "os"
	"net/http"

    "github.com/gin-gonic/gin"
    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type product struct {
    Product_id  int   `json:"id"`
    Name  string  `json:"title"`
    Spec string  `json:"artist"`
    Catg    string
    Price  float64 `json:"price"`
}

type inventory struct {
    Id     int   `json:"id"`
    Quantity int 
}



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
    //     Catg: "bottle"
    //     Price  150.0
    // })



    router.GET("/products", getProducts)
//  router.GET("/products/:id", getProductsByCatg)
    router.GET("/products/:id", getProductByID)
    router.POST("/products", postProducts)

    router.Run("localhost:8080")
}

// getProducts responds with the list of all products as JSON.
func getProducts(c *gin.Context) {

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
        if err := rows.Scan(&pdt.Product_id, &pdt.Name, &pdt.Spec, &pdt.Catg, &pdt.Price); err != nil {
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
func getProductByID(c *gin.Context) {
    id,_ := strconv.Atoi(c.Param("id"))

	var pdt product

    row := db.QueryRow("SELECT * FROM product WHERE product_id = ?", id)
    if err := row.Scan(&pdt.Product_id, &pdt.Name, &pdt.Spec, &pdt.Catg, &pdt.Price); err!=nil{
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


// postProducts adds an product from JSON received in the request body.
func postProducts(c *gin.Context) {
    var newProduct product

    // Call BindJSON to bind the received JSON to
    // newProduct.
    if err := c.BindJSON(&newProduct); err != nil {
        return
    }

    result, err := db.Exec("INSERT INTO product (product_id, name, spec, catg, price) VALUES (?, ?, ?, ?, ?)", newProduct.Product_id, newProduct.Name, newProduct.Spec, newProduct.Catg, newProduct.Price)
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

// &pdt.Product_id, &pdt.Name, &pdt.Spec, &pdt.Catg, &pdt.Price