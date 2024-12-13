package models

import "time"

type UserRole struct {
	UserID    uint       `gorm:"primaryKey" json:"user_id,omitempty"`
	RoleID    uint       `gorm:"primaryKey" json:"role_id,omitempty"`
	User      User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Role      Role       `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE" json:"role,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at,omitempty"`
	CreatedBy string     `json:"created_by,omitempty"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	UpdatedBy string     `json:"updated_by,omitempty"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"` // Soft delete with timestamp
	DeletedBy string     `json:"deleted_by,omitempty"`
}