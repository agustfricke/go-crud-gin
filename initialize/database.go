package initialize

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DataBase() {
    var err error
    dsn := "host=localhost user=agust password=agust dbname=some-postgres port=5432"
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Error connecting database")
    } 
}
