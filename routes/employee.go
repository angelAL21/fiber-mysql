package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

//serializer of employee
type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    int     `json:"age"`
}

var DB *gorm.DB

//endpoints

func GetEmployees(c *fiber.Ctx) error {
	var employees []Employee
	DB.Find(&employees)
	return c.JSON(&employees)
}

func GetEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee Employee
	DB.Find(&employee, id)
	return c.JSON(&employee)
}

func CreateEmployee(c *fiber.Ctx) error {
	employee := new(Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&employee)
	return c.JSON(&employee)
}
