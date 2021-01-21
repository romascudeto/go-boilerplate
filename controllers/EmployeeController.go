package controllers

import (
	"fmt"
	"go-project/helpers"
	"go-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetEmployee ... Get all employees
func GetEmployee(c *gin.Context) {
	var employee []models.Employee
	err := models.GetAllEmployee(&employee)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		meta := map[string]interface{}{}
		response := helpers.NewResponse(meta, employee)
		c.JSON(http.StatusOK, response)
	}
}

//CreateEmployee ... Create Employee
func CreateEmployee(c *gin.Context) {
	var employee models.Employee
	c.BindJSON(&employee)
	err := models.CreateEmployee(&employee)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

//GetEmployeeByID ... Get the employee by id
func GetEmployeeByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var employee models.Employee
	err := models.GetEmployeeByID(&employee, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		meta := map[string]interface{}{}
		response := helpers.NewResponse(meta, employee)
		c.JSON(http.StatusOK, response)
	}
}

//UpdateEmployee ... Update the employee information
func UpdateEmployee(c *gin.Context) {
	var employee models.Employee
	id := c.Params.ByName("id")
	err := models.GetEmployeeByID(&employee, id)
	if err != nil {
		c.JSON(http.StatusNotFound, employee)
	}
	c.BindJSON(&employee)
	err = models.UpdateEmployee(&employee, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

//DeleteEmployee ... Delete the employee
func DeleteEmployee(c *gin.Context) {
	var employee models.Employee
	id := c.Params.ByName("id")
	err := models.DeleteEmployee(&employee, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
