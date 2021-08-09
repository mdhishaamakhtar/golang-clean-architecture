package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
	"github.com/mdhishaamakhtar/learnFiber/api/views"
	"github.com/spf13/viper"
	"log"
)

func Protected() func(*fiber.Ctx) error {
	key, _ := viper.Get("JWT_SECRET").(string)
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(key),
		ErrorHandler: jwtError,
		ContextKey:   "Token",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Missing or malformed JWT",
		})

	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid or expired JWT",
		})
	}
}

func ValidateAndGetClaims(c *fiber.Ctx) (map[string]interface{}, error) {
	token, ok := c.Locals("Token").(*jwt.Token)
	if !ok {
		log.Println(token)
		return nil, views.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Println(claims)
		return nil, views.ErrInvalidToken
	}

	if claims.Valid() != nil {
		return nil, views.ErrInvalidToken
	}
	return claims, nil
}
