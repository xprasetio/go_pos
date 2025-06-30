package middleware

import (
	"net/http"
	"strconv"
	"strings"

	user_domain "pos-go/internal/user"
	"pos-go/pkg/jwt"

	"github.com/labstack/echo/v4"
)

type JWTMiddleware struct {
	jwtService  jwt.JWTService
	userService user_domain.UserService
}

func NewJWTMiddleware(jwtService jwt.JWTService, userService user_domain.UserService) *JWTMiddleware {
	return &JWTMiddleware{
		jwtService:  jwtService,
		userService: userService,
	}
}

func (m *JWTMiddleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Authorization header required",
			})
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid authorization header format",
			})
		}

		// Extract the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		userID, email, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid or expired token",
			})
		}

		// Ambil role_id dari userService menggunakan userID
		user, err := m.userService.FindByID(userID)
		if err != nil || user == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User not found"})
		}

		// Set user information in context
		c.Set("user_id", strconv.FormatUint(uint64(userID), 10))
		c.Set("user_email", email)
		c.Set("role_id", user.RoleID)

		return next(c)
	}
}

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Log request
		req := c.Request()
		method := req.Method
		uri := req.RequestURI
		remoteAddr := req.RemoteAddr

		// Log before processing
		c.Logger().Infof("Request: %s %s from %s", method, uri, remoteAddr)

		// Process request
		err := next(c)

		// Log response
		status := c.Response().Status
		c.Logger().Infof("Response: %s %s - %d", method, uri, status)

		return err
	}
}
