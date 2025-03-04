package user_controller

import (
	"api/database"
	"api/models/blog_model"
	"api/request"
	"api/responses"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetAllBlog(ctx *gin.Context) {
	// Mendeklarasikan variabel blog yang akan menampung data user
	blogs := new([]models.Blog)

	// Menarik data dari tabel blog di database
	err := database.DB.Table("blogs").Find(&blogs).Error

	// Jika terjadi error dalam menarik data dari database
	if err != nil {
		// Mengembalikan respons error jika gagal mengambil data
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	// Mengembalikan data siswa dalam bentuk JSON dengan status 200
	ctx.JSON(200, gin.H{
		"data": blogs, // Mengembalikan data siswa yang berhasil diambil
	})
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

	ctx.JSON(200, gin.H{
		"message" : "data transmitted",
		"data": blog,
	})
}


func Store(ctx *gin.Context) {
	blogReq := request.BlogRequest{
		Title:       ctx.PostForm("title"),
		Description: ctx.PostForm("description"),
		Category:    ctx.PostForm("category"),
		Content:     ctx.PostForm("content"),
	}

	// Cek apakah title sudah digunakan
	var existingBlog models.Blog
	if err := database.DB.Where("title = ?", blogReq.Title).First(&existingBlog).Error; err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Title already used."})
		return
	}

	// Menangani file header_photo
	var filename string
	file, err := ctx.FormFile("header_photo")
	if err == nil { // Jika file ada
		ext := filepath.Ext(file.Filename)
		filename = fmt.Sprintf("%s%s", blogReq.Title, ext)
		savePath := fmt.Sprintf("../web/public/uploads/%s", filename)

		if err := ctx.SaveUploadedFile(file, savePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to upload file."})
			return
		}
	} else {
		filename = ""
	}

	// Simpan ke database
	blog := models.Blog{
		Title:       &blogReq.Title,
		Description: &blogReq.Description,
		Category:    &blogReq.Category,
		Content:     &blogReq.Content,
		HeaderPhoto: &filename,
	}

	if err := database.DB.Create(&blog).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't create data."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data Successfully Created",
		"data":    blog,
	})
}
