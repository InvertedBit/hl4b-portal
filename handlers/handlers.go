package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/InvertedBit/hl4b-portal/config"
	"github.com/InvertedBit/hl4b-portal/database"
	"github.com/InvertedBit/hl4b-portal/middleware"
	"github.com/InvertedBit/hl4b-portal/models"
	"github.com/InvertedBit/hl4b-portal/templates"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"maragu.dev/gomponents"
)

type Handlers struct {
	store  *session.Store
	config *config.Config
}

func NewHandlers(store *session.Store, cfg *config.Config) *Handlers {
	return &Handlers{
		store:  store,
		config: cfg,
	}
}

// renderComponent renders a gomponents node to HTML
func (h *Handlers) renderComponent(c *fiber.Ctx, component gomponents.Node) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return component.Render(c.Response().BodyWriter())
}

// Home handles the home page
func (h *Handlers) Home(c *fiber.Ctx) error {
	sess, err := h.store.Get(c)
	if err != nil {
		return h.renderComponent(c, templates.HomePage("", false, false))
	}

	userID := sess.Get("user_id")
	if userID == nil {
		return h.renderComponent(c, templates.HomePage("", false, false))
	}

	username := sess.Get("username")
	isAdmin := sess.Get("is_admin")

	usernameStr := ""
	if username != nil {
		usernameStr = username.(string)
	}

	isAdminBool := false
	if isAdmin != nil {
		isAdminBool = isAdmin.(bool)
	}

	return h.renderComponent(c, templates.HomePage(usernameStr, isAdminBool, true))
}

// LoginPage shows the login form
func (h *Handlers) LoginPage(c *fiber.Ctx) error {
	errorMsg := c.Query("error", "")
	return h.renderComponent(c, templates.LoginPage(errorMsg))
}

// Login handles login authentication
func (h *Handlers) Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return c.Redirect("/login?error=Email and password are required")
	}

	// For demo purposes, we'll check if the user exists in auth.users
	// In production, you would verify against the actual auth system
	var authUser models.AuthUser
	result := database.DB.Where("email = ?", email).First(&authUser)

	if result.Error != nil {
		// For demo: Create the user if they don't exist
		// In production, this would be handled by your auth system
		authUser = models.AuthUser{
			ID:        uuid.New(),
			Email:     email,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		database.DB.Create(&authUser)
	}

	// Check if portal user exists, create if not
	var portalUser models.PortalUser
	result = database.DB.Where("user_id = ?", authUser.ID).First(&portalUser)

	if result.Error != nil {
		// Create portal user
		username := email[:len(email)-len(filepath.Ext(email))] // Use email prefix as username
		portalUser = models.PortalUser{
			UserID:   authUser.ID,
			Username: username,
			IsAdmin:  false, // First user could be admin, or set manually
		}
		database.DB.Create(&portalUser)
	}

	// Create session
	sess, err := h.store.Get(c)
	if err != nil {
		return c.Redirect("/login?error=Session error")
	}

	sess.Set("user_id", authUser.ID.String())
	sess.Set("username", portalUser.Username)
	sess.Set("is_admin", portalUser.IsAdmin)

	if err := sess.Save(); err != nil {
		return c.Redirect("/login?error=Failed to create session")
	}

	return c.Redirect("/user/dashboard")
}

// Logout handles user logout
func (h *Handlers) Logout(c *fiber.Ctx) error {
	sess, err := h.store.Get(c)
	if err == nil {
		sess.Destroy()
	}
	return c.Redirect("/")
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a password with a hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// UserDashboard shows the user dashboard
func (h *Handlers) UserDashboard(c *fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return err
	}

	username := c.Locals("username").(string)
	isAdmin := middleware.IsAdmin(c)

	// Get upload count
	var uploadCount int64
	database.DB.Model(&models.Upload{}).Where("user_id = ?", userID).Count(&uploadCount)

	return h.renderComponent(c, templates.UserDashboard(username, isAdmin, int(uploadCount)))
}

// UserUploads shows the user's uploads
func (h *Handlers) UserUploads(c *fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return err
	}

	username := c.Locals("username").(string)
	isAdmin := middleware.IsAdmin(c)

	var uploads []models.Upload
	database.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&uploads)

	return h.renderComponent(c, templates.UserUploads(username, isAdmin, uploads))
}

// UploadFile handles file uploads
func (h *Handlers) UploadFile(c *fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return err
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("No file uploaded")
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	newFilename := fmt.Sprintf("%s_%d%s", uuid.New().String(), time.Now().Unix(), ext)
	filePath := filepath.Join(h.config.UploadsDir, newFilename)

	// Save file
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save file")
	}

	// Create database record
	upload := models.Upload{
		UserID:       userID,
		Filename:     newFilename,
		OriginalName: file.Filename,
		FileSize:     file.Size,
		MimeType:     file.Header.Get("Content-Type"),
		FilePath:     filePath,
	}

	if err := database.DB.Create(&upload).Error; err != nil {
		os.Remove(filePath) // Clean up file if database insert fails
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save upload record")
	}

	// Return the upload item component
	return h.renderComponent(c, templates.UploadItem(upload))
}

// DeleteUpload handles file deletion
func (h *Handlers) DeleteUpload(c *fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return err
	}

	uploadID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid upload ID")
	}

	var upload models.Upload
	result := database.DB.Where("id = ? AND user_id = ?", uploadID, userID).First(&upload)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("Upload not found")
	}

	// Delete file from disk
	os.Remove(upload.FilePath)

	// Delete from database
	database.DB.Delete(&upload)

	return c.SendString("")
}
