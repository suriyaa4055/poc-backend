package initializers

import (
	"fmt"
	"testBench/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {
	var err error
	// DB_URL := "host=host.docker.internal user=root password=root dbname=assignment1 port=5432 sslmode=disable"
	DB_URL := "host=localhost user=postgres password=Suriyaa dbname=testBench port=5432 sslmode=disable"
	// DB_URL := "host=172.17.0.2 user=root password=root dbname=assignment1 port=5432 sslmode=disable"

	// host.docker.internal
	// DB_URL := "host=host.docker.internal user=root password=root dbname=assignment1 port=5432 sslmode=disable"
	//dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to db")
	}
	DB.AutoMigrate(&models.TestBenchTable{})
	DB.AutoMigrate(&models.Details{})

	return DB
}
