package main

import (
	"fmt"
	"log"

	db "github.com/arnavmahajan630/go-hrms-api/DB"
	"github.com/gofiber/fiber/v2"
)



func main() {
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	 app := fiber.New()
	// app.Get("/employee")
	// app.Post("/employee")
	// app.Put("/employee/:id")
	// app.Delete("/employee/:id")
    

	fmt.Println("Server Started on Port 8080")
	err = app.Listen(":8080")
	if err != nil {
		fmt.Println(err)
	}

}