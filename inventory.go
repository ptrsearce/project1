package main

import (
	"net/http"
	"strconv"

	"fmt"

	"github.com/gin-gonic/gin"
)

type inventory struct {
	Inventory_id int `json:"id"`
	Quantity     int
}

func GetInventoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var ivt inventory

	row := db.QueryRow("SELECT * FROM inventory WHERE inventory_id = ?", id)
	if err := row.Scan(&ivt.Inventory_id, &ivt.Quantity); err != nil {
		// if err == sql.ErrNoRows {
		//    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "inventory not found"})
		//    fmt.Errorf("inventorysById %d: no such inventory", id)
		//    return
		// }
		// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "inventory not found"})
		// fmt.Errorf("inventorysById %d: %v", id, err)
		// return
	}
	c.IndentedJSON(http.StatusOK, ivt)

}

// getInventorys responds with the list of all inventorys as JSON.
func GetInventorys(c *gin.Context) {

	var inventorys []inventory

	rows, err := db.Query("SELECT * FROM inventory")
	if err != nil {
		fmt.Errorf("inventorys : %v", err)
		return
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var ivt inventory
		if err := rows.Scan(&ivt.Inventory_id, &ivt.Quantity); err != nil {
			fmt.Errorf("inventorys : %v", err)
			return
		}
		inventorys = append(inventorys, ivt)
	}
	if err := rows.Err(); err != nil {
		fmt.Errorf("inventorys : %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, inventorys)
}

// getInventoryByID locates the inventory whose ID value matches the id
// parameter sent by the client, then returns that inventory as a response.

// postInventorys adds an inventory from JSON received in the request body.
func PostInventorys(c *gin.Context) {
	var newInventory inventory

	// Call BindJSON to bind the received JSON to
	// newInventory.
	if err := c.BindJSON(&newInventory); err != nil {
		return
	}

	result, err := db.Exec("INSERT INTO inventory (inventory_id, quantity) VALUES (?, ?)", newInventory.Inventory_id, newInventory.Quantity)
	// Add the new inventory to the slice.
	if err != nil {
		fmt.Errorf("addInventory: %v", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Errorf("addInventory: %v", err)
		return
	}
	fmt.Print(id)
	c.IndentedJSON(http.StatusCreated, newInventory)
}

// &ivt.Inventory_id, &ivt.Name, &ivt.Spec, &ivt.Catg_id, &ivt.Price
