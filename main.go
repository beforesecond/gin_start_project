package main

import (
	"log"
	"test-gin/databases"
	"test-gin/routers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func Migrate(db *gorm.DB) {
	//users.AutoMigrate()
	//db.AutoMigrate(&articles.ArticleModel{})
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := databases.GetDB()
	Migrate(db)
	defer db.Close()
	r := gin.Default()
	v1 := r.Group("/api")
	routers.Router(v1.Group("/v1"))
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
