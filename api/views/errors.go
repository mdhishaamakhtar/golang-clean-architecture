package views

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/mdhishaamakhtar/learnFiber/pkg"
)

type ErrView struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//noinspection ALL
var (
	ErrMethodNotAllowed = errors.New("Error: Method is not allowed")
	ErrInvalidToken     = errors.New("Error: Invalid Authorization token")
	ErrUserExists       = errors.New("Error: User already exists")
	ErrFile             = errors.New("Error: Something wrong with file")
	ErrUpload           = errors.New("Error: Upload failed")
)

var ErrHTTPStatusMap = map[string]int{
	pkg.ErrNotFound.Error():     fiber.StatusNotFound,
	pkg.ErrInvalidSlug.Error():  fiber.StatusBadRequest,
	pkg.ErrExists.Error():       fiber.StatusConflict,
	pkg.ErrNoContent.Error():    fiber.StatusNotFound,
	pkg.ErrDatabase.Error():     fiber.StatusInternalServerError,
	pkg.ErrUnauthorized.Error(): fiber.StatusUnauthorized,
	pkg.ErrForbidden.Error():    fiber.StatusForbidden,
	pkg.ErrEmail.Error():        fiber.StatusBadRequest,
	pkg.ErrPassword.Error():     fiber.StatusBadRequest,
	ErrMethodNotAllowed.Error(): fiber.StatusMethodNotAllowed,
	ErrInvalidToken.Error():     fiber.StatusBadRequest,
	ErrUserExists.Error():       fiber.StatusConflict,
	ErrFile.Error():             fiber.StatusBadRequest,
}

func Wrap(err error, c *fiber.Ctx) error {
	msg := err.Error()
	code := ErrHTTPStatusMap[msg]

	if code == 0 {
		code = fiber.StatusInternalServerError
	}

	errView := ErrView{
		Error:   true,
		Message: msg,
		Status:  code,
	}

	return c.Status(code).JSON(errView)
}
