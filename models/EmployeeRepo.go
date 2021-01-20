package models

import (
	"fmt"
	Config "go-project/config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllEmployees Fetch all employee data
func GetAllEmployee(employee *[]Employee) (err error) {
	if err = Config.DB.Find(employee).Error; err != nil {
		return err
	}
	return nil
}

//CreateEmployee ... Insert New data
func CreateEmployee(employee *Employee) (err error) {
	if err = Config.DB.Create(employee).Error; err != nil {
		return err
	}
	return nil
}

//GetEmployeeByID ... Fetch only one employee by Id
func GetEmployeeByID(employee *Employee, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(employee).Error; err != nil {
		return err
	}
	return nil
}

//UpdateEmployee ... Update employee
func UpdateEmployee(employee *Employee, id string) (err error) {
	fmt.Println(employee)
	Config.DB.Save(employee)
	return nil
}

//DeleteEmployee ... Delete employee
func DeleteEmployee(employee *Employee, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(employee)
	return nil
}
