package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// order structure contains serialization information
type order struct {
	Id               string `json:"id"`
	Created          string `json:"created"`
	Status           string `json:"status"`
	Customer         string `json:"customer"`
	Sku              string `json:"sku"`
	Photo            string `json:"photo"`
	Condition        string `json:"condition"`
	Size             string `json:"size"`
	Type             string `json:"type"`
	Origin_address   string `json:"origin_address"`
	Shipping_address string `json:"shipping_address"`
}

// Array of orders to store
var allOrders []order

// Function to seed db with initial data from orders.json
func seedData() (orders []order) {
	// Attempt to read orders.json
	content, err := ioutil.ReadFile("orders.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	// Convert json to go struct
	json.Unmarshal(content, &orders)

	return orders
}

// Start of web endpoints
// Accepts gin context (contains all information related to a request)
// Returns response of OK with all orders
// /orders
func getOrders(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, allOrders)
}

// Create new Order
// /create
func createOrder(c *gin.Context) {
	var newOrder order

	// Attempt to bind json to order struct returns an error
	if err := c.BindJSON(&newOrder); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Order creation error."})
		return
	}

	allOrders = append(allOrders, newOrder)

	c.IndentedJSON(http.StatusCreated, newOrder)
}

// Update Order
// /update
func updateOrder(c *gin.Context) {
	var newOrder order
	if err := c.BindJSON(&newOrder); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Update binding error."})
		return
	}
	var update bool = false

	// check if order exists with ID
	for idx, x := range allOrders {
		if x.Id == newOrder.Id {
			x = newOrder
			allOrders[idx] = x
			update = true
			break
		}
	}

	if update {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Order updated."})
		return
	} else {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}
}

// check if order id exists
func getOrderById(id string) (*order, error) {
	for i, o := range allOrders {
		if o.Id == id {
			return &allOrders[i], nil
		}
	}
	return nil, errors.New("Order not found")
}

func main() {
	allOrders = seedData()

	fmt.Println(allOrders)

	// Web api
	router := gin.Default()

	// setup cors
	router.Use(cors.Default())

	// Endpoint for all orders
	router.GET("/orders", getOrders)

	// Endpoint for creating order
	router.POST("/create", createOrder)

	// Endpoint for update orders
	router.PATCH("/update", updateOrder)

	// Run server on router
	router.Run("localhost:8081")
}
