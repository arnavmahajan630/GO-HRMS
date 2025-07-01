package helpers

import (
	db "github.com/arnavmahajan630/go-hrms-api/DB"
	"github.com/arnavmahajan630/go-hrms-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// private
func getEmployeeCollection() *mongo.Collection {
	return db.Mg.DB.Collection("employees")
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

func AddEmployee(c * fiber.Ctx) error{
    ctx, cancel := db.TimeOutCkeck()
    defer cancel()

    var employee models.Employee
    if err := c.BodyParser(&employee); err != nil {
      return c.Status(400).JSON(fiber.Map{"error": "Invalid Json format"})
   }

   employee.ID = primitive.NewObjectID()
   collection := getEmployeeCollection()
   _, err := collection.InsertOne(ctx, employee)
   if err != nil {
      return c.Status(500).JSON(fiber.Map{"error":"Insertion Failed"})
   }
   
   return c.Status(200).JSON(fiber.Map{"message": "Inserted Successfully"})
   
}

func DeleteEmployee(c * fiber.Ctx) error{
   ctx , cancel := db.TimeOutCkeck()
   defer cancel()

   id := c.Params("id")
   objId, err := primitive.ObjectIDFromHex(id)
   if err != nil {
      c.Status(400).JSON(fiber.Map{"error": "Invalid Id"})
   }
   collection := getEmployeeCollection()

   q := bson.M{"_id": objId}

   res, err := collection.DeleteOne(ctx, q)
   if err != nil {
      return c.Status(500).JSON(fiber.Map{"error": "Delete failed"})

   }
   if res.DeletedCount == 0 {
      return c.Status(500).JSON(fiber.Map{"error": "Emplyee not found"}) 
   }

   return c.Status(200).JSON(fiber.Map{"message": "Deleted successfully"})

   
}

func EditEmployee(c * fiber.Ctx) error{
   ctx , cancel := db.TimeOutCkeck()
   defer cancel()

   id := c.Params("id")
   objId , err := primitive.ObjectIDFromHex(id)
   if err != nil {
      return c.Status(400).JSON(fiber.Map{"error":"Invalid Id"})
   }
   var employee models.Employee
   if err := c.BodyParser(&employee); err != nil {
      return c.Status(400).JSON(fiber.Map{"error": "Invalid Json Type"})
   }

   collection := getEmployeeCollection()

   update := bson.D {
      {
         Key: "$set", Value: bson.D{
            {Key: "name", Value: employee.Name},
            {Key: "age", Value: employee.Age},
            {Key: "salary", Value: employee.Salary},
         }},
   }
    res , err := collection.UpdateByID(ctx, objId, update)
    if err != nil || res.MatchedCount == 0 {
      return c.Status(500).JSON(fiber.Map{"error":"Update failed. Id not found"})
    }

    return c.JSON(fiber.Map{"message": "update succeeded"})

}
