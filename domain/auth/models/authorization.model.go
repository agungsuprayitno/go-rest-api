package models

import (
	"time"

	"github.com/google/uuid"
)

type Authorization struct {
	ID        				uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	ClientID     			string      `gorm:"uniqueIndex;" json:"client_id,omitempty"`
	ClientSecret   			string      `json:"client_secret,omitempty"`
	AuthorizationCode		string      `json:"authorization_code,omitempty"`
	Scope     				string      `json:"scope,omitempty"`
	RequestIP     			string      `json:"request_ip,omitempty"`
	CreatedAt 				time.Time   `json:"created_at,omitempty"`
	ExpiredAt 				time.Time   `json:"expired_at,omitempty"`
}
