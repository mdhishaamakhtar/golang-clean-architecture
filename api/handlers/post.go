package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdhishaamakhtar/learnFiber/api/middlewares"
	"github.com/mdhishaamakhtar/learnFiber/api/views"
	"github.com/mdhishaamakhtar/learnFiber/pkg/models"
	"github.com/mdhishaamakhtar/learnFiber/pkg/post"
)

func AddPost(svc post.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		a, err := middlewares.ValidateAndGetClaims(c)
		if err != nil {
			return views.Wrap(err, c)
		}
		p := &models.Post{}
		if err := c.BodyParser(p); err != nil {
			return views.Wrap(err, c)
		}
		p.UserID = a["id"].(string)
		err = svc.AddPost(p)
		if err != nil {
			return views.Wrap(err, c)
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":   false,
			"message": "Post Created",
		})
	}
}

func GetPosts(svc post.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		a, err := middlewares.ValidateAndGetClaims(c)
		if err != nil {
			return views.Wrap(err, c)
		}
		p, _ := svc.GetAllPosts(a["id"].(string))
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   false,
			"message": "Posts Fetched",
			"posts":   p,
		})
	}
}

func MakePostHandler(app *fiber.App, svc post.Service) {
	postGroup := app.Group("/api/v1/post")
	postGroup.Post("/add", middlewares.Protected(), AddPost(svc))
	postGroup.Get("/get", middlewares.Protected(), GetPosts(svc))
}
