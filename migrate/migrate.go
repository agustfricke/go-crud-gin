package main

import (
	"github.com/agustfricke/go-crud-gin/initialize"
	"github.com/agustfricke/go-crud-gin/models"
)


func init() {
    initialize.LoadEnv()
    initialize.DataBase()
}

func main() {
    initialize.DB.AutoMigrate(&models.Post{})
}

