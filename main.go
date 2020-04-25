package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/karanbirsingh7/go-fiber-sample/book"
	"github.com/karanbirsingh7/go-fiber-sample/database"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello World")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Put("/api/v1/book/:id", book.UpdateBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Successfully connected to database")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database successfully migrated")
}

func main() {
	// INIT app
	app := fiber.New()

	// Database connection open and close
	initDatabase()
	defer database.DBConn.Close()

	// Routes
	setupRoutes(app)

	app.Listen(3000)
}
