package controller

import (
	"go-blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Blog struct {
	DB *gorm.DB
}

type postForm struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type blogResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (b *Blog) FindAll(ctx *gin.Context) {
	blogs := []models.Blog{}

	if err := b.DB.Order("id desc").Find(&blogs).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var serializedBlog []blogResponse
	copier.Copy(&serializedBlog, &blogs)
	ctx.JSON(http.StatusCreated, gin.H{"post": serializedBlog})
}

func (b *Blog) FindOne(ctx *gin.Context) {
	blog, err := b.findPostByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var serializedBlog blogResponse
	copier.Copy(&serializedBlog, &blog)
	ctx.JSON(http.StatusOK, gin.H{"post": serializedBlog})
}

func (b *Blog) Update(ctx *gin.Context) {
	var form postForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	blog, err := b.findPostByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	copier.Copy(&blog, &form)

	if err := b.DB.Save(&blog).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "updated..."})
}

func (b *Blog) Delete(ctx *gin.Context) {
	blog, err := b.findPostByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := b.DB.Delete(&blog).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (b *Blog) Create(ctx *gin.Context) {
	var form postForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var blog models.Blog

	copier.Copy(&blog, &form)

	if err := b.DB.Create(&blog).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success"})
}

func (b *Blog) findPostByID(ctx *gin.Context) (*models.Blog, error) {
	var blog models.Blog
	id := ctx.Param("id")
	if err := b.DB.First(&blog, id).Error; err != nil {
		return nil, err
	}

	return &blog, nil

}
