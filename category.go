package main

import (
	"strconv"
	"net/http"
    
    "fmt"

    "github.com/gin-gonic/gin"
)

type category struct {
	Category_id int `json:"id"`
	Name     string
}

func GetCategoryByID(c *gin.Context) {
    id,_ := strconv.Atoi(c.Param("id"))

	var ctg category

    row := db.QueryRow("SELECT * FROM category WHERE category_id = ?", id)
    if err := row.Scan(&ctg.Category_id, &ctg.Name); err!=nil{
        // if err == sql.ErrNoRows {
        //    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "category not found"})
        //    fmt.Errorf("categorysById %d: no such category", id) 
        //    return
        // }
        // c.IndentedJSON(http.StatusNotFound, gin.H{"message": "category not found"})
        // fmt.Errorf("categorysById %d: %v", id, err)
        // return
	}
	c.IndentedJSON(http.StatusOK, ctg)

}

// getCategorys responds with the list of all categorys as JSON.
func GetCategorys(c *gin.Context) {

	var categorys []category

	rows, err := db.Query("SELECT * FROM category")
	if err != nil {
		fmt.Errorf("categorys : %v", err)
		return
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var ctg category
		if err := rows.Scan(&ctg.Category_id, &ctg.Name); err != nil {
			fmt.Errorf("categorys : %v", err)
			return
		}
		categorys = append(categorys, ctg)
	}
	if err := rows.Err(); err != nil {
		fmt.Errorf("categorys : %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, categorys)
}

// getCategoryByID locates the category whose ID value matches the id
// parameter sent by the client, then returns that category as a response.

// postCategorys adds an category from JSON received in the request body.
func PostCategorys(c *gin.Context) {
	var newCategory category

	// Call BindJSON to bind the received JSON to
	// newCategory.
	if err := c.BindJSON(&newCategory); err != nil {
		return
	}

	result, err := db.Exec("INSERT INTO category (category_id, name) VALUES (?, ?)", newCategory.Category_id, newCategory.Name)
	// Add the new category to the slice.
	if err != nil {
		fmt.Errorf("addCategory: %v", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Errorf("addCategory: %v", err)
		return
	}
	fmt.Print(id)
	c.IndentedJSON(http.StatusCreated, newCategory)
}

// &ctg.Category_id, &ctg.Name, &ctg.Spec, &ctg.Catg_id, &ctg.Price