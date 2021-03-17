package main

import (
	"GoRestAPIDemo/database"
	"GoRestAPIDemo/models"
	"GoRestAPIDemo/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("Movies.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	// Migrate the schemas
	database.DBConn.AutoMigrate(&models.Movie{})
}

func setupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"message": "You are at the endpoint 😉",
		})
	})

	api := app.Group("/api")

	api.Get("", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"message": "You are at the endpoint 😉",
		})
	})

	routes.MovieRoute(api.Group("/movies"))
}

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	initDatabase()

	setupRoutes(app)

	app.Listen(":8080")
}
