package main

import (
        "fmt"
        "github.com/gofiber/fiber/v2"
        "github.com/fiber/book"
        "github.com/fiber/database"
        "github.com/jinzhu/gorm"
        _"github.com/jinzhu/gorm/dialects/postgres"
       )


func initDatabase() {
  var err error
  database.DBConn, err = gorm.Open("postgres", #postegres_details)
  if err != nil {
    panic("Failed to Connect to Database")
  }
  fmt.Println("Connection to database established.")

  database.DBConn.AutoMigrate(&book.Book{})
  fmt.Println("Database Migrated")

}


func helloWorld(c *fiber.Ctx) error{
  return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App){
  app.Get("/api/v1/book", book.GetBooks)
  app.Get("/api/v1/book/:id", book.GetBook)
  app.Post("api/v1/book", book.NewBook)
  app.Delete("api/v1/book/:id", book.DeleteBook)
}


func main(){
  app := fiber.New()

  initDatabase()
  defer database.DBConn.Close()

  setupRoutes(app)

  app.Listen(":3000")
}
