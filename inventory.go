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
type inventory struct {
	Inventory_id int `json:"inventory_id"`
	Quantity     int `json:"quantity"`
}


func CreateInventoryTable(db *sql.DB) error {  
    query :=`CREATE TABLE IF NOT EXISTS inventory (
		inventory_id int unsigned NOT NULL,
		quantity int DEFAULT NULL,
		PRIMARY KEY (inventory_id),
		UNIQUE KEY inventory_id_UNIQUE (inventory_id),
		CONSTRAINT inventory_id FOREIGN KEY (inventory_id) REFERENCES inventory (inventory_id)
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci`

    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    res, err := db.ExecContext(ctx, query)
    if err != nil {
        log.Printf("Error %s when creating inventory table", err)
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


func GetInventoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var ivt inventory

	row := db.QueryRow("SELECT * FROM inventory WHERE inventory_id = ?", id)
	if err := row.Scan(&ivt.Inventory_id, &ivt.Quantity); err != nil {
		if err == sql.ErrNoRows {
		   c.IndentedJSON(http.StatusNotFound, gin.H{"message": "inventory not found"})
		   fmt.Errorf("inventorysById %d: no such inventory", id)
		   return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "inventory not found"})
		fmt.Errorf("inventorysById %d: %v", id, err)
		return
	}
	c.IndentedJSON(http.StatusOK, ivt)

}

// getInventorys responds with the list of all inventorys as JSON.
func GetInventory(c *gin.Context) {

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
func PostInventory(c *gin.Context) {
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

func DeleteInventoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var ivt inventory

	row := db.QueryRow("SELECT * FROM inventory WHERE inventory_id = ?", id)
	if err := row.Scan(&ivt.Inventory_id, &ivt.Quantity); err != nil {
		if err == sql.ErrNoRows {
		   c.IndentedJSON(http.StatusNotFound, gin.H{"message": "inventory not found"})
		   fmt.Errorf("inventorysById %d: no such inventory", id)
		   return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "inventory not found"})
		fmt.Errorf("inventorysById %d: %v", id, err)
		return
	}
	db.QueryRow("DELETE FROM inventory WHERE inventory_id = ?", id)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "inventory is deleted"})

}

// &ivt.Inventory_id, &ivt.Quantity, &ivt.Spec, &ivt.Catg_id, &ivt.Price
func UpdateInventoryByID(c *gin.Context) {
	var newInventory inventory
	id, _ := strconv.Atoi(c.Param("id"))
	// Call BindJSON to bind the received JSON to
	// newInventory.
	if err := c.BindJSON(&newInventory); err != nil {
		return
	}

	if(newInventory.Inventory_id !=0){db.Exec("UPDATE inventory SET inventory_id= ? WHERE inventory_id = ?",newInventory.Inventory_id, id)}
	if(newInventory.Quantity !=0){db.Exec("UPDATE inventory SET quantity= ? WHERE inventory_id = ?",newInventory.Quantity, id)}
	// Add the new inventory to the slice.
	// if err != nil {
	// 	fmt.Errorf("addInventory: %v", err)
	// 	return
	// }
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	fmt.Errorf("addInventory: %v", err)
	// 	return
	// }
	fmt.Print(id)
	c.IndentedJSON(http.StatusCreated, newInventory)

}