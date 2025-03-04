package migration

import (
    "log"
    "api/database"
    blogModel "api/models/blog"
    userModel "api/models/user"
)

func MigrateDatabase() {
	database.ConnectDatabase()

	err := database.DB.AutoMigrate(
		&blogModel.Blog{},
		&userModel.User{}, 
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration successful!")
}