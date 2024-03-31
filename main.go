// main.go
package main

import (
	importdata "assignment/Importdata"
	"assignment/handlers"
	"assignment/redisOperations"
	"assignment/sqlOperations"

	"github.com/gin-gonic/gin"
)

func init() {
	// Initialize MySQL and Redis connections
	err := sqlOperations.InitializeMySQL()
	if err != nil {
		panic(err)
	}
	err = redisOperations.InitializeRedis()
	if err != nil {
		panic(err)
	}

	// Import data from Excel
	filePath := "./uploads/Sample_Employee_Data.xlsx"
	err = importdata.ImportDataFromExcel(filePath)
	if err != nil {
		panic(err)
	}
}

func main() {
	// Initialize Gin router
	router := gin.Default()

	// API routes
	router.GET("/employees", handlers.GetAllEmployeesHandler)
	router.GET("/employee/:id", handlers.GetEmployeeHandler)
	router.PUT("/employee/:id", handlers.UpdateEmployeeHandler)
	router.DELETE("/employee/:id", handlers.DeleteEmployeeHandler)
	router.GET("/redis/employees", handlers.GetAllEmployeesFromRedisHandler)
	router.GET("/redis/employee/:id", handlers.GetEmployeeFromRedisHandler)

	// Start server
	router.Run(":8080")
}
