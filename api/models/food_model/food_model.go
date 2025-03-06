package models

type Food struct {
	ID       *int    `json:"id"`
	Name     *string `gorm:"type:varchar(255);not null" json:"name"`
	Calorie  *float64 `gorm:"type:float;not null" json:"calorie"`
	Carbohydrate *float64 `gorm:"type:float;not null" json:"carbohydrate"`
	Category    *string `gorm:"type:varchar(255);not null" json:"category"`
	Cholesterol *float64 `gorm:"type:float;not null" json:"cholesterol"`
	Fat *float64 `gorm:"type:float;not null" json:"fat"`
	Image *string `gorm:"type:varchar(255);not null" json:"image"`
	MonounsaturatedFat *float64 `gorm:"type:float;not null" json:"monounsaturated_fat"`
	PolyunsaturatedFat *float64 `gorm:"type:float;not null" json:"polyunsaturated_fat"`
	Portion *float64 `gorm:"type:float;not null" json:"portion"`
	Potassium *float64 `gorm:"type:float;not null" json:"potassium"`
	Protein *float64 `gorm:"type:float;not null" json:"protein"`
	SaturatedFat *float64 `gorm:"type:float;not null" json:"saturated_fat"`
	Sodium *float64 `gorm:"type:float;not null" json:"sodium"`
	Sugar *float64 `gorm:"type:float;not null" json:"sugar"`
	Unit *string `gorm:"type:varchar(255);not null" json:"unit"`
}