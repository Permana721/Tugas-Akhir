package models

import (
    "time"
)

type Blog struct {
    ID          *uint      `gorm:"primaryKey" json:"id"`
    Title       *string    `gorm:"type:varchar(255);not null" json:"title"`
    Description *string    `gorm:"type:varchar(255);not null" json:"description"`
    Category    *string    `gorm:"type:varchar(255);not null" json:"category"`
    Content     *string    `gorm:"type:text;not null" json:"content"`
    HeaderPhoto *string    `gorm:"type:varchar(255);not null" json:"header_photo"`
    Photo       *string    `gorm:"type:varchar(255)" json:"photo"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
}