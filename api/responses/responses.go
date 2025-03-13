package responses

import (
    "time"
)

type BlogResponse struct {
	ID          *uint      `json:"id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Category    *string    `json:"category"`
	Content     *string    `json:"content"`
	Image 		*string    `json:"image"`
}

type UserResponse struct {	
	ID     		*int    	`json:"id"`
	Name 		*string 	`json:"name"`
	Email     	*string  	`json:"email"`
	Password    *string  	`json:"password"`
	Age     	*int    	`json:"age"`
}

type FoodResponse struct {
	ID          		*uint		`json:"id"`
	Name        		*string 	`json:"name"`
	Calorie  			*float64	`json:"calorie"`
	Carbohydrate 		*float64 	`json:"carbohydrate"`
	Category    		*string 	`json:"category"`
	Cholesterol 		*float64 	`json:"cholesterol"`
	Fat 				*float64 	`json:"fat"`
	Fiber 				*float64 	`json:"fiber"`
	Image 				*string 	`json:"image"`
	MonounsaturatedFat 	*float64 	`json:"monounsaturated_fat"`
	PolyunsaturatedFat 	*float64 	`json:"polyunsaturated_fat"`
	Portion	 			*float64 	`json:"portion"`
	Potassium 			*float64 	`json:"potassium"`
	Protein 			*float64 	`json:"protein"`
	SaturatedFat 		*float64 	`json:"saturated_fat"`
	Sodium 				*float64 	`json:"sodium"`
	Sugar 				*float64 	`json:"sugar"`
	Unit 				*string 	`json:"unit"`
}

type SportResponse struct {
	ID          		*uint      `json:"id"`
    Name       			*string    `json:"name"`
    Intensity 			*string    `json:"intensity"`
    CalorieBurn    		*float64   `json:"calorie_burn"`
    AvarageDuration   	*int   	   `son:"avarage_duration"`
    TargetMuscle      	*string    `json:"target_muscle"`
    SuitableLevel       *string    `json:"suitable_level"`
    Description       	*string    `json:"description"`
    SportType       	*string    `json:"sport_type"`
    CreatedAt   		time.Time  `json:"created_at"`
    UpdatedAt   		time.Time  `json:"updated_at"`
}