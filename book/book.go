package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/karanbirsingh7/go-fiber-sample/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	var books []Book
	db := database.DBConn

	// Find all book and append value to books array
	db.Find(&books)

	// Send real books object in response
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	var book Book
	db := database.DBConn
	db.Find(&book, id)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	var book Book
	db := database.DBConn

	db.First(&book, id)

	// IF NO book is found
	if book.Title == "" {
		c.Status(404).Send("No book found with given ID")
		return
	}

	db.Delete(&book, book.ID)
	c.JSON(book.Title + " deleted")
}

func NewBook(c *fiber.Ctx) {
	book := new(Book)
	db := database.DBConn

	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&book)
	c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	bookRequest := new(Book)
	bookOrg := new(Book)

	// Book data from request
	if err := c.BodyParser(bookRequest); err != nil {
		c.Status(503).Send(err)
		return
	}

	// Check if bookID return any valid books
	db.First(&bookOrg, id)
	if bookOrg.Title == "" {
		c.Status(404).Send("No book found with given ID")
		return
	}

	db.Model(&Book{}).Update(&bookRequest)
	c.JSON(&bookRequest)
}
