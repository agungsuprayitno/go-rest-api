package models

import (
	"go-rest-postgres/domain/users/models"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Title     string      `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Content   string      `gorm:"not null" json:"content,omitempty"`
	Image     string      `gorm:"not null" json:"image,omitempty"`
	User      uuid.UUID   `gorm:"not null" json:"user,omitempty"`
	UserData  models.User `gorm:"foreignKey:User" json:"user_data,omitempty"`
	CreatedAt time.Time   `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time   `gorm:"not null" json:"updated_at,omitempty"`
}
