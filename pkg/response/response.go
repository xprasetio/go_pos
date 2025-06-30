package response

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponseHttp(c echo.Context, code int, message string, data interface{}) error {
	resp := Response{
		Message: message,
		Data:    data,
	}
	return c.JSON(code, resp)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash memverifikasi password dengan hash yang tersimpan
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
