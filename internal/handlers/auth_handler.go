package handlers

import (
	"sample-app/internal/models"
	"sample-app/internal/services"
	"sample-app/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userService *services.UserService
}

func NewAuthHandler(userService *services.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

func (h *AuthHandler) Signup(c *fiber.Ctx) error {
	request := new(models.SignupRequest)
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := request.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "validation error",
			"fields":  err,
		})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.ErrInternalServerError)
	}

	user := new(models.User)
	user.Username = request.Username
	user.Email = request.Email
	user.Password = string(password)

	err = h.userService.Create(user)
	if err != nil {
		if utils.IsDuplicateEntryError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "username or email already exists",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.ErrInternalServerError)
	}

	return c.Status(fiber.StatusAccepted).JSON(user.ToUserResponse())
}

func (h *AuthHandler) Signin(c *fiber.Ctx) error {
	request := new(models.SigninRequest)
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := request.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "validation error",
			"fields":  err,
		})
	}

	user := new(models.User)
	user.Email = request.Email
	user.Username = request.Username

	err = h.userService.Retrieve(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.ErrUnauthorized)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.ErrUnauthorized)
	}

	return c.Status(fiber.StatusAccepted).JSON(user.ToUserResponse())
}

func (h AuthHandler) SetRoute(a *fiber.App) {
	group := a.Group("/auth")
	group.Post("/signup", h.Signup)
	group.Post("/signin", h.Signin)
}
