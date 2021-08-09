package handlers

import (
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/mdhishaamakhtar/learnFiber/api/middlewares"
	"github.com/mdhishaamakhtar/learnFiber/api/views"
	"github.com/mdhishaamakhtar/learnFiber/pkg/models"
	"github.com/mdhishaamakhtar/learnFiber/pkg/user"
	"github.com/spf13/viper"
	"log"
)

func Register(svc user.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		u := &models.User{}
		if err := c.BodyParser(u); err != nil {
			return views.Wrap(err, c)
		}
		exists, err := svc.GetUserDetailsByEmail(u.Email)
		if err != nil {
			return views.Wrap(err, c)
		}
		if exists.ID != "" {
			return views.Wrap(views.ErrUserExists, c)
		}
		err = svc.AddUser(u)
		if err != nil {
			return views.Wrap(err, c)
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":   false,
			"message": "User Created",
		})
	}
}

func Login(svc user.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		u := &models.User{}
		if err := c.BodyParser(u); err != nil {
			return views.Wrap(err, c)
		}
		us, err := svc.Login(u.Email, u.Password)
		if err != nil {
			return views.Wrap(err, c)
		}
		us.Password = ""
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    us.ID,
			"email": us.Email,
		})
		secret, ok := viper.Get("JWT_SECRET").(string)
		if !ok {
			log.Panic(fmt.Errorf("jwt secret not found"))
		}
		tokenString, err := token.SignedString([]byte(secret))
		if err != nil {
			return views.Wrap(err, c)
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   false,
			"message": "user logged in",
			"details": us,
			"token":   tokenString,
		})
	}
}

func GetDetails(svc user.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		a, err := middlewares.ValidateAndGetClaims(c)
		if err != nil {
			return views.Wrap(err, c)
		}
		us, _ := svc.GetUserDetailsById(a["id"].(string))
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   false,
			"message": "User Details Fetched",
			"details": us,
		})
	}
}

func MakeUserHandler(app *fiber.App, svc user.Service) {
	userGroup := app.Group("/api/v1/user")
	userGroup.Post("/register", Register(svc))
	userGroup.Post("/login", Login(svc))
	userGroup.Get("/detail", middlewares.Protected(), GetDetails(svc))
}
