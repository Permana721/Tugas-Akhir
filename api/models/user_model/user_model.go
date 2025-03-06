package models

type User struct {
	ID       *int    `json:"id"`
	Name     *string `gorm:"type:varchar(255);not null" json:"name"`
	Email    *string `gorm:"type:varchar(255);not null" json:"email"`
	Password *string `gorm:"type:varchar(255);not null" json:"password"`
	Age      *int    `gorm:"type:integer" json:"age"`
}