package models

import (
	"time"

	"github.com/google/uuid"
)

type Merchant struct {
	ID        				uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	MerchantID     			string      `gorm:"uniqueIndex;not null" json:"merchant_id,omitempty"`
	MerchantName   			string      `json:"merchant_name,omitempty"`
	PhoneNumber				*string     `json:"phone_number,omitempty"`
	Email     				string      `json:"email,omitempty"`
	ClientID     			string      `json:"client_id,omitempty"`
	ClientSecret			string      `json:"client_secret,omitempty"`
	PIC						string      `json:"pic,omitempty"`
	PICEmail				string      `json:"pic_email,omitempty"`
	PICPhoneNumber			*string     `json:"pic_phone_number,omitempty"`
	CreatedAt 				time.Time	`json:"created_at,omitempty"`
	CreatedBy 				string		`json:"created_by,omitempty"`
	ModifiedAt 				*time.Time	`json:"modified_at,omitempty"`
	ModifiedBy 				*string		`json:"modified_by,omitempty"`
	DeletedAt 				*time.Time	`json:"deleted_at,omitempty"`
}
