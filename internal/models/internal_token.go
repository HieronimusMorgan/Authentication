package models

import (
	"gorm.io/gorm"
	"time"
)

type InternalToken struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	ResourceID uint           `gorm:"not null" json:"resource_id"`
	Token      string         `gorm:"not null" json:"token"`
	Expired    bool           `gorm:"not null" json:"expired"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at,omitempty"`
	CreatedBy  string         `json:"created_by,omitempty"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	UpdatedBy  string         `json:"updated_by,omitempty"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	DeletedBy  string         `json:"deleted_by,omitempty"`
}
