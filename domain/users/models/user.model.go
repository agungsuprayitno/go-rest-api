package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Role      string    `gorm:"type:varchar(255);not null" json:"role"`
	Provider  string    `gorm:"not null" json:"provider"`
	Photo     string    `gorm:"not null" json:"photo"`
	Verified  bool      `gorm:"not null" json:"verified"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}
