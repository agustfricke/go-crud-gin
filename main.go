package main

import (
	"github.com/agustfricke/go-crud-gin/controllers"
	"github.com/agustfricke/go-crud-gin/initialize"
	"github.com/gin-gonic/gin"
)

func init() {
    initialize.LoadEnv()
    initialize.DataBase()
}

func main() {
	r := gin.Default()

	r.POST("/", controllers.CreatePost)
	r.GET("/", controllers.GetPosts)
    r.GET("/:id", controllers.GetSoloPost)
    r.PUT("/:id", controllers.PutPost)
    r.DELETE("/:id", controllers.DeletePost)

	r.Run() 
}
