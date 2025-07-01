package helpers

import (
	db "github.com/arnavmahajan630/go-hrms-api/DB"
	"github.com/arnavmahajan630/go-hrms-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// private
func getEmployeeCollection() *mongo.Collection {
	return db.Mg.DB.Collection("employees")
}

func AddEmployees(c * fiber.Ctx) error{
   return nil
}

// exported
func GetEmployees(c * fiber.Ctx) error{
   ctx , cancel := db.TimeOutCkeck()
   defer cancel()
   collection := getEmployeeCollection()
   querry := bson.D{}
   cursor, err := collection.Find(ctx, querry)
   if err != nil {return c.Status(500).JSON(fiber.Map{"error": "DB not found"})}
   defer cursor.Close(ctx)

   var employees []models.Employee
   if err := cursor.All(ctx, &employees); err != nil {
      return c.Status(500).JSON(fiber.Map{"error": "Internal server error. Unable to decode"})
   }

   return c.Status(200).JSON(employees)
   

}

func EditEmployee(c * fiber.Ctx) error{
    return nil
}

func DeleteEmployee(c * fiber.Ctx) error{
   return nil
}