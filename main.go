package main

import (
	"fmt"
	"log"

	db "github.com/arnavmahajan630/go-hrms-api/DB"
	"github.com/arnavmahajan630/go-hrms-api/handlers"
	"github.com/gofiber/fiber/v2"
)



func main() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	// connected to Mongodb

	app := fiber.New()
	handlers.RoutesInit(app)

	// initialized routes
	
	fmt.Println("Server started successfully on port 8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}

	// server started and running

	
	

}