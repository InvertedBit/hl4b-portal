package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
)

// AuthRequired middleware ensures the user is authenticated
func AuthRequired(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Redirect("/login")
		}

		userID := sess.Get("user_id")
		if userID == nil {
			return c.Redirect("/login")
		}

		// Store user_id in locals for handlers to use
		c.Locals("user_id", userID)
		c.Locals("is_admin", sess.Get("is_admin"))
		c.Locals("username", sess.Get("username"))

		return c.Next()
	}
}

// AdminRequired middleware ensures the user is an administrator
func AdminRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isAdmin := c.Locals("is_admin")
		
		if isAdmin == nil || !isAdmin.(bool) {
			return c.Status(fiber.StatusForbidden).SendString("Access denied: Admin privileges required")
		}

		return c.Next()
	}
}

// GetUserID retrieves the user ID from the context
func GetUserID(c *fiber.Ctx) (uuid.UUID, error) {
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	return userID, nil
}

// IsAdmin checks if the current user is an admin
func IsAdmin(c *fiber.Ctx) bool {
	isAdmin := c.Locals("is_admin")
	if isAdmin == nil {
		return false
	}
	return isAdmin.(bool)
}
