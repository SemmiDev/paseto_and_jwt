package main

import (
	"github.com/gofiber/fiber/v2"
)

func routes(app fiber.Router) {
	router := app.Group("/api/v1")

	router.Post("/login", LoginHandler)
	router.Post("/register", RegisterHandler)
}

func LoginHandler(c *fiber.Ctx) error {
	var loginRequest LoginUserRequest
	err := c.BodyParser(&loginRequest)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := loginRequest.Check()

	if err == ErrUserNotFound {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error": ErrUserNotFound,
		})
	}

	if err == ErrUsernameAndPasswordIsWrong {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error": ErrUsernameAndPasswordIsWrong,
		})
	}

	response := NewLoginUserResponse(NewUserResponse(user))
	return c.JSON(fiber.Map{
		"data": response,
	})
}

func RegisterHandler(c *fiber.Ctx) error {
	var user CreateUserRequest
	err := c.BodyParser(&user)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	newUserResponse := user.createNewUser()
	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"user": newUserResponse,
	})
}