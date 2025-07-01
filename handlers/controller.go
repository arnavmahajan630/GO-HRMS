package handlers

import (
	"github.com/arnavmahajan630/go-hrms-api/helpers"
	"github.com/gofiber/fiber/v2"
)

func RoutesInit(app *fiber.App) {
	// CRUD Endpoints
	app.Get("/employee", helpers.GetEmployees)
	app.Post("/employee", helpers.AddEmployees)
	app.Put("/employee/:id", helpers.EditEmployee)
	app.Delete("/employee/:id", helpers.DeleteEmployee)

}
