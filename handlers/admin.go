package handlers

import (
	"os"
	"strconv"

	"github.com/InvertedBit/hl4b-portal/database"
	"github.com/InvertedBit/hl4b-portal/models"
	"github.com/InvertedBit/hl4b-portal/templates"
	"github.com/gofiber/fiber/v2"
)

// AdminDashboard shows the admin dashboard
func (h *Handlers) AdminDashboard(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	// Get statistics
	var userCount int64
	database.DB.Model(&models.PortalUser{}).Count(&userCount)

	var uploadCount int64
	database.DB.Model(&models.Upload{}).Count(&uploadCount)

	return h.renderComponent(c, templates.AdminDashboard(username, int(userCount), int(uploadCount)))
}

// AdminUsers shows all users for management
func (h *Handlers) AdminUsers(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	var users []models.PortalUser
	database.DB.Order("created_at DESC").Find(&users)

	return h.renderComponent(c, templates.AdminUsers(username, users))
}

// UpdateUserRole updates a user's admin status
func (h *Handlers) UpdateUserRole(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	isAdmin := c.FormValue("is_admin") == "true"

	var user models.PortalUser
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	user.IsAdmin = isAdmin
	database.DB.Save(&user)

	// Return updated row
	return h.renderComponent(c, templates.AdminUserRow(user))
}

// DeleteUser deletes a user
func (h *Handlers) DeleteUser(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	var user models.PortalUser
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	// Delete user's uploads first
	var uploads []models.Upload
	database.DB.Where("user_id = ?", user.UserID).Find(&uploads)
	for _, upload := range uploads {
		os.Remove(upload.FilePath)
	}
	database.DB.Where("user_id = ?", user.UserID).Delete(&models.Upload{})

	// Delete portal user
	database.DB.Delete(&user)

	return c.SendString("")
}

// AdminUploads shows all uploads for management
func (h *Handlers) AdminUploads(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	var uploads []models.Upload
	database.DB.Order("created_at DESC").Find(&uploads)

	// Get user information for display
	var portalUsers []models.PortalUser
	database.DB.Find(&portalUsers)

	userMap := make(map[string]string)
	for _, user := range portalUsers {
		userMap[user.UserID.String()] = user.Username
	}

	return h.renderComponent(c, templates.AdminUploads(username, uploads, userMap))
}

// AdminDeleteUpload deletes any upload (admin privilege)
func (h *Handlers) AdminDeleteUpload(c *fiber.Ctx) error {
	uploadID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid upload ID")
	}

	var upload models.Upload
	if err := database.DB.First(&upload, uploadID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Upload not found")
	}

	// Delete file from disk
	os.Remove(upload.FilePath)

	// Delete from database
	database.DB.Delete(&upload)

	return c.SendString("")
}
