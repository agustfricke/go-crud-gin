package controllers

import (
	"github.com/agustfricke/go-crud-gin/initialize"
	"github.com/agustfricke/go-crud-gin/models"
	"github.com/gin-gonic/gin"
)

func DeletePost(c *gin.Context) {
    id := c.Param("id") 

    initialize.DB.Delete(&models.Post{}, id)

    c.Status(204)
}

func PutPost(c *gin.Context) {
    id := c.Param("id") 

    var PostData struct {
        Title string
        Body string
    }

    c.Bind(&PostData)

    var post models.Post
    initialize.DB.First(&post, id)

    initialize.DB.Model(&post).Updates(models.Post{
        Title: PostData.Title,
        Body: PostData.Body,
    })

    c.JSON(200, gin.H{
        "post": post,
    }) 
}

func GetSoloPost(c *gin.Context) {
    id := c.Param("id")
    var post models.Post
    initialize.DB.First(&post, id)

    c.JSON(200, gin.H{
       "solo post": post,
   }) 
    
}

func GetPosts(c *gin.Context) {
    var posts []models.Post
    initialize.DB.Find(&posts)

    c.JSON(200, gin.H{
       "posts": posts,
   }) 
}

func CreatePost(c *gin.Context) {

    var PostData struct {
        Title string
        Body string
    }

    c.Bind(&PostData)

    post := models.Post{Title: PostData.Title, Body:PostData.Body}

    result := initialize.DB.Create(&post) 

    if result.Error != nil {
        c.Status(400)
        return 
    }

   c.JSON(200, gin.H{
       "post": post,
   }) 
}

