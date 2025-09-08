package middleware

import (
	"scs-operator/pkg/errors"
	"scs-operator/pkg/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

func (mw *MiddlewareManager) JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return errors.NewUnauthorizedError("missing token")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return errors.NewUnauthorizedError("invalid token format")
		}

		tokenString := parts[1]
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			return errors.NewUnauthorizedError("invalid or expired token")
		}

		// Store claims in context
		c.Set("user_id", claims.UserID)

		return next(c)
	}
}
