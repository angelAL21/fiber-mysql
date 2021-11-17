package main

import (
	"fmt"
	"log"

	"github.com/angelAL21/fiber-mysql/database"
	"github.com/angelAL21/fiber-mysql/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("hi")

	// Connect with database
	database.InitialMigration()

	// Create a Fiber app
	app := fiber.New()
	log.Fatal(app.Listen(":3000"))

	//endpoints
	//createEmployee
	app.Post("/api/employee", routes.CreateEmployee)
	//GetAllEmployees
	app.Get("/api/employees", routes.GetEmployees)
	//GetEmployeeById
	app.Get("/api/employee/:id", routes.GetEmployee)
}
