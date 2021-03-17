package routes

import (
	"GoRestAPIDemo/controllers"

	"github.com/gofiber/fiber/v2"
)

func MovieRoute(route fiber.Router) {
	route.Get("", controllers.GetMovies)
	route.Post("", controllers.CreateMovie)
	route.Get(":title", controllers.GetMovie)
	route.Put("update", controllers.UpdateMovie)
	route.Delete("/delete/:id", controllers.DeleteMovie)
}
