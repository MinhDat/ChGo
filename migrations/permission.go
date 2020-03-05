package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigratePermission() {
	db := db.GetDB()
	if !db.HasTable(&models.Permission{}) {
		err := db.CreateTable(&models.Permission{})
		if err != nil {
			log.Println("Permission table already exists")
		}
	}
}
