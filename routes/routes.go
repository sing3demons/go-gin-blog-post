package routes

import (
	"go-blog/config"
	"go-blog/controller"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	v1 := r.Group("/api/v1")

	blogController := controller.Blog{DB: db}

	blogGroup := v1.Group("/posts")
	{
		blogGroup.GET("", blogController.FindAll)
		blogGroup.POST("", blogController.Create)
		blogGroup.GET("/:id",blogController.FindOne)
		blogGroup.PUT("/:id",blogController.Update)
		blogGroup.DELETE("/:id",blogController.Delete)
	}
}
