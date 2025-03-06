package blog_controller

import (
	"api/database"
	models "api/models/blog_model"
	"api/responses"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetAllBlog(ctx *gin.Context) {
	blogs := new([]models.Blog)
	err := database.DB.Table("blogs").Find(&blogs).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, blogs)
}

func GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	blog := new(responses.BlogResponse)

	errDb := database.DB.Table("blogs").Where("id = ?", id).Find(&blog).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	if blog.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}

	ctx.JSON(200, blog)
}

func Store(ctx *gin.Context) {
	title :=       ctx.PostForm("title")
	description := ctx.PostForm("description")
	category :=    ctx.PostForm("category")
	content :=     ctx.PostForm("content")

	var existingBlog models.Blog
	if err := database.DB.Where("title = ?", title).First(&existingBlog).Error; err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Title already used."})
		return
	}

	var filename string
	file, err := ctx.FormFile("image")
	if err == nil { 
		ext := filepath.Ext(file.Filename)
		filename = fmt.Sprintf("%s%s", title, ext)
		savePath := fmt.Sprintf("../web/public/uploads/%s", filename)

		if err := ctx.SaveUploadedFile(file, savePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to upload file."})
			return
		}
	} else {
		filename = ""
	}

	blog := models.Blog{
		Title:       &title,
		Description: &description,
		Category:    &category,
		Content:     &content,
		Image: &filename,
	}

	if err := database.DB.Create(&blog).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't create data."})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	blog := new(models.Blog)
	blogTitleExist := new(models.Blog)

	title :=       ctx.PostForm("title")
	description := ctx.PostForm("description")
	category :=    ctx.PostForm("category")
	content :=     ctx.PostForm("content")

	if title == "" || description == "" || category == "" || content == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Semua field harus diisi!",
		})
		return
	}

	errDb := database.DB.Table("blogs").Where("id = ?", id).Find(&blog).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if blog.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}

	errblogTitleExist := database.DB.Table("blogs").Where("title = ?", title).Find(&blogTitleExist).Error
	if errblogTitleExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if blogTitleExist.Title != nil && *blog.ID != *blogTitleExist.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Title already used",
		})
		return
	}

	file, err := ctx.FormFile("image")
    var filename string
    if err == nil {
        ext := filepath.Ext(file.Filename)
        filename = fmt.Sprintf("%s%s", title, ext)

        savePath := fmt.Sprintf("../web/public/uploads/%s", filename)

        if err := ctx.SaveUploadedFile(file, savePath); err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to upload file."})
            return
        }
    } else {
        filename = *blog.Image
    }

	blog.Title = &title
	blog.Description = &description
	blog.Category = &category
	blog.Content = &content
	blog.Image = &filename

	errUpdate := database.DB.Table("blogs").Where("id = ?", id).Updates(&blog).Error
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't update data",
		})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func Delete(ctx *gin.Context) {
    id := ctx.Param("id")
    blog := new(models.Blog)

    result := database.DB.Table("blogs").Where("id = ?", id).Find(&blog)
    if result.Error != nil {
        ctx.JSON(500, gin.H{
            "message": "internal server error",
        })
        return
    }
    if result.RowsAffected == 0 {
        ctx.JSON(404, gin.H{
            "message": "data not found",
        })
        return
    }

    errDb := database.DB.Table("blogs").Where("id = ?", id).Delete(&models.Blog{}).Error
    if errDb != nil {
        ctx.JSON(500, gin.H{
            "message": "internal server error",
            "error": errDb.Error(),
        })
        return
    }

    ctx.JSON(200, gin.H{
        "message": "data deleted successfully",
    })
}