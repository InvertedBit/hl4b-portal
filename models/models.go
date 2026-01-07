package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AuthUser represents the existing auth.users table in PostgreSQL
// This table is managed externally and should not be migrated
type AuthUser struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName specifies the table name for AuthUser
func (AuthUser) TableName() string {
	return "auth.users"
}

// PortalUser represents custom user data for the portal
// This links to the auth.users table via UserID
type PortalUser struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"` // References auth.users.id
	Username  string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	IsAdmin   bool      `gorm:"default:false;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Upload represents a file uploaded by a user
type Upload struct {
	ID           uint      `gorm:"primaryKey"`
	UserID       uuid.UUID `gorm:"type:uuid;not null;index"` // References auth.users.id
	Filename     string    `gorm:"type:varchar(255);not null"`
	OriginalName string    `gorm:"type:varchar(255);not null"`
	FileSize     int64     `gorm:"not null"`
	MimeType     string    `gorm:"type:varchar(100)"`
	FilePath     string    `gorm:"type:varchar(500);not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
