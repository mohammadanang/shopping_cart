package middleware

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadanang/shopping-cart/pkg/token"
)

func (m *Middleware) ValidateBearer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get(authorizationHeader)
		if len(authHeader) == 0 {
			return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
				"message": "authorization header is not provided",
				"status":  "error",
				"code":    fiber.ErrUnauthorized.Code,
			})
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
				"message": "invalid authorization header format",
				"status":  "error",
				"code":    fiber.ErrUnauthorized.Code,
			})
		}

		authType := strings.ToLower(fields[0])
		if authType != authorizationBearerType {
			return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
				"message": fmt.Errorf("unsupported authorization type %s", authType).Error(),
				"status":  "error",
				"code":    fiber.ErrUnauthorized.Code,
			})
		}

		accessToken := fields[1]
		payload, err := m.tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
				"message": err.Error(),
				"status":  "error",
				"code":    fiber.ErrUnauthorized.Code,
			})
		}

		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
				"message": err.Error(),
				"status":  "error",
				"code":    fiber.ErrUnauthorized.Code,
			})
		}

		c.Locals(authorizationPayload, string(payloadBytes))
		return c.Next()
	}
}

func (m *Middleware) ExtractBearer(c *fiber.Ctx) (payload *token.Payload) {
	stringPayload := c.Locals(authorizationPayload).(string)
	err := json.Unmarshal([]byte(stringPayload), &payload)
	if err != nil {
		fmt.Printf("error extracting token payload %s", err.Error())
		return nil
	}

	return payload
}
