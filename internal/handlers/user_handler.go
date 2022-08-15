package handlers

import (
	"sample-app/internal/models"
	"sample-app/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (h UserHandler) Users(c *fiber.Ctx) error {
	users, err := h.userService.All()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}
	var response []models.UserResponse
	for _, user := range users {
		response = append(response, user.ToResponse())
	}
	return c.JSON(response)
}

func (h UserHandler) Retrieve(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}
	user, err := h.userService.RetrieveById(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}
	return c.JSON(user.ToResponse())
}

func (h UserHandler) Create(c *fiber.Ctx) error {
	createRequest := new(models.User)
	if err := c.BodyParser(&createRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	notValid := createRequest.Validate()
	if notValid != nil {
		return c.Status(fiber.StatusBadRequest).JSON(notValid)
	}

	err := h.userService.Create(createRequest)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}
	return c.Status(fiber.StatusAccepted).JSON(createRequest.ToResponse())
}

func (h UserHandler) Update(c *fiber.Ctx) error {
	updateRequest := new(models.User)
	if err := c.BodyParser(&updateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	user, err := h.userService.RetrieveById(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	err = h.userService.Update(&user, *updateRequest)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(err)
	}

	return c.Status(fiber.StatusAccepted).JSON(user.ToResponse())
}

func (h UserHandler) Delete(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	err = h.userService.Delete(userId)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(err)
	}

	return c.Status(fiber.StatusAccepted).Send(nil)
}

func (h UserHandler) SetRoute(a *fiber.App) {
	group := a.Group("/users")
	group.Get("/", h.Users)
	group.Post("/", h.Create)
	group.Get("/:id", h.Retrieve)
	group.Put("/:id", h.Update)
	group.Delete("/:id", h.Delete)
}
