package main

import (
	"log"
	"os"

	"github.com/InvertedBit/hl4b-portal/config"
	"github.com/InvertedBit/hl4b-portal/database"
	"github.com/InvertedBit/hl4b-portal/handlers"
	"github.com/InvertedBit/hl4b-portal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	database.Connect(cfg.DatabaseURL)
	database.AutoMigrate()

	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll(cfg.UploadsDir, 0755); err != nil {
		log.Fatal("Failed to create uploads directory:", err)
	}

	// Create session store
	store := session.New()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "HL4B Portal",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Static files
	app.Static("/static", "./static")
	app.Static("/uploads", cfg.UploadsDir)

	// Initialize handlers
	h := handlers.NewHandlers(store, cfg)

	// Public routes
	app.Get("/", h.Home)
	app.Get("/login", h.LoginPage)
	app.Post("/login", h.Login)
	app.Post("/logout", h.Logout)

	// Protected user routes
	user := app.Group("/user", middleware.AuthRequired(store))
	user.Get("/dashboard", h.UserDashboard)
	user.Get("/uploads", h.UserUploads)
	user.Post("/upload", h.UploadFile)
	user.Delete("/upload/:id", h.DeleteUpload)

	// Admin routes
	admin := app.Group("/admin", middleware.AuthRequired(store), middleware.AdminRequired())
	admin.Get("/dashboard", h.AdminDashboard)
	admin.Get("/users", h.AdminUsers)
	admin.Post("/users/:id/role", h.UpdateUserRole)
	admin.Delete("/users/:id", h.DeleteUser)
	admin.Get("/uploads", h.AdminUploads)
	admin.Delete("/uploads/:id", h.AdminDeleteUpload)

	// Start server
	port := cfg.Port
	if port == "" {
		port = "3000"
	}
	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
