package databases

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDB() *gorm.DB {
	//db := gorm.close()
	db, _ := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=postgres")
	return db
}
