package models

import (
    "time"
)

type Sport struct {
    ID          		*uint      `gorm:"primaryKey" json:"id"`
    Name       			*string    `gorm:"type:varchar(100);not null" json:"name"`
    Intensity 			*string    `gorm:"type:varchar(50);not null" json:"intensity"`
    CalorieBurn    		*float64   `gorm:"type:float;not null" json:"calorie_burn"`
    AvarageDuration   	*int   	   `gorm:"type:int;not null" json:"avarage_duration"`
    TargetMuscle      	*string    `gorm:"type:varchar(150);not null" json:"target_muscle"`
    SuitableLevel       *string    `gorm:"type:varchar(100);not null" json:"suitable_level"`
    Description       	*string    `gorm:"type:text;not null" json:"description"`
    SportType       	*string    `gorm:"type:varchar(150);not null" json:"sport_type"`
    CreatedAt   		time.Time  `json:"created_at"`
    UpdatedAt   		time.Time  `json:"updated_at"`
}