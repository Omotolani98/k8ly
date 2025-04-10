// services/logly/models/log.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	ID        uint           `gorm:"primaryKey"`
	Message   string         `gorm:"type:text;not null"`
	Service   string         `gorm:"type:varchar(100);not null"`
	Level     string         `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
