package main

import (
	"distributed_database_server/config"
	auth "distributed_database_server/internal/auth/entity"
	"fmt"
	// todo "distributed_database_server/internal/todo/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// Connect to database
	c := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", c.PostgreSQL.Host, c.PostgreSQL.User, c.PostgreSQL.Password, c.PostgreSQL.DBName, c.PostgreSQL.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connect to database successfully")

	fmt.Println("Run migrate ...")

	// Migrate the schema
	db.AutoMigrate(&auth.User{})
	//db.AutoMigrate(&todo.Todo{})

	fmt.Println("Migrate successfully")
}
