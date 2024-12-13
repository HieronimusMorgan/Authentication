package models

import "time"

type Role struct {
	RoleID      uint       `gorm:"primaryKey" json:"role_id,omitempty"`
	Name        string     `gorm:"unique;not null" json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at,omitempty"`
	CreatedBy   string     `json:"created_by,omitempty"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	UpdatedBy   string     `json:"updated_by,omitempty"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"` // Soft delete with timestamp
	DeletedBy   string     `json:"deleted_by,omitempty"`
}
