package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Car struct {
	ID    string `json:"id"`
	Model string `json:"model"`
	Color string `json:"color"`
}

var cars []Car

func main() {
	cars = []Car{
		{ID: "1", Model: "Audi", Color: "Blue"},
		{ID: "2", Model: "BMW", Color: "Green"},
	}

	router := gin.Default()
	router.GET("cars", GetCars)
	router.GET("cars/:id", GetCarById)
	router.POST("cars", createCar)
	router.DELETE("cars/:id", deleteCar)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func deleteCar(c *gin.Context) {
	for index, car := range cars {
		if car.ID == c.Param("id") {
			cars = append(cars[:index], cars[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Car has been deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Car not found"})
}

func createCar(c *gin.Context) {
	var newCar Car
	if err := c.BindJSON(&newCar); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	cars = append(cars, newCar)
	c.IndentedJSON(http.StatusCreated, newCar)
}

func GetCarById(c *gin.Context) {
	for _, car := range cars {
		if car.ID == c.Param("id") {
			c.IndentedJSON(200, car)
			return
		}
	}
}

func GetCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"cars": cars})
}
