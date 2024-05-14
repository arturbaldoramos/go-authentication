package middleware

import (
	"github.com/arturbaldoramos/go-authentication/pkg/template"
	"github.com/arturbaldoramos/go-authentication/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"os"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Cookies("token")
	if tokenString == "" {
		component := template.ErrorPage(
			"Unauthorized",
			"You must log in to access this page",
			"401")
		return utils.Render(c, component)
	}

	// Parse JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Malformed token",
				})
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Token is expired or not yet valid",
				})
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Couldn't handle this token",
				})
			}
		}
	}

	// Check if token is valid
	if token.Valid {
		// Continue chain
		return c.Next()
	}

	// Continue chain
	return c.Next()
}
