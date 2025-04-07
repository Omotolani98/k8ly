// models/request.go
package models

import "gorm.io/gorm"

type LoggedRequest struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt int64          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Method  string `gorm:"type:varchar(10)" json:"method"`
	URL     string `gorm:"type:text" json:"url"`
	Headers string `gorm:"type:text" json:"headers"`
	Body    string `gorm:"type:text" json:"body"`
}
