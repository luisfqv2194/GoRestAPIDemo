package controllers

import (
	. "GoRestAPIDemo/database"
	. "GoRestAPIDemo/models"

	"github.com/gofiber/fiber/v2"
)

func GetMovies(c *fiber.Ctx) error {
	var movies []Movie
	DBConn.Find(&movies)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"movies": movies,
		},
	})
}

func GetMovie(c *fiber.Ctx) error {
	title := "%" + c.Params("title") + "%"
	var movie Movie
	if err := DBConn.First(&movie, "Title LIKE ?", title).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"err":     err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"movie": movie,
		},
	})

}

func CreateMovie(c *fiber.Ctx) error {

	movie := new(Movie)
	if err := c.BodyParser(movie); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	DBConn.Create(&movie)
	return c.JSON(movie)
}

func UpdateMovie(c *fiber.Ctx) error {
	updatedMovie := new(Movie)
	if err := c.BodyParser(updatedMovie); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	id := updatedMovie.ID
	var originalMovie Movie
	if err := DBConn.First(&originalMovie, id).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"err":     err.Error(),
		})
	}

	DBConn.Model(&originalMovie).Updates(updatedMovie)
	return c.JSON(updatedMovie)
}

func DeleteMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	var movie Movie
	if err := DBConn.First(&movie, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"err":     err.Error(),
		})
	}
	DBConn.Delete(&movie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Movie deleted",
	})

}
