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
    Author      *string    `gorm:"type:varchar(255);not null" json:"author"`
    Image       *string    `gorm:"type:varchar(255);not null" json:"image"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
}