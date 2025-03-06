package migration

import (
    "log"
    "api/database"
    blogModel "api/models/blog_model"
	foodModel "api/models/food_model"
    userModel "api/models/user_model"
)

func MigrateDatabase() {
	database.ConnectDatabase()

	err := database.DB.AutoMigrate(
		&blogModel.Blog{},
		&foodModel.Food{},
		&userModel.User{}, 
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration successful!")
}