package food_controller

import (
	"api/database"
	models "api/models/food_model"
	"api/responses"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllFood(ctx *gin.Context) {
	foods := new([]models.Food)
	err := database.DB.Table("foods").Find(&foods).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, foods)
}

func GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	food := new(responses.FoodResponse)

	errDb := database.DB.Table("foods").Where("id = ?", id).Find(&food).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	if food.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}

	ctx.JSON(200, food)
}

func parseFloat(ctx *gin.Context, field string) (float64, bool) {
	valueStr := ctx.PostForm(field)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%s harus berupa angka!", field)})
		return 0, false
	}
	return value, true
}

func Store(ctx *gin.Context) {
	name := ctx.PostForm("name")
	category := ctx.PostForm("category")
	unit := ctx.PostForm("unit")

	fields := []string{"calorie", "carbohydrate", "cholesterol", "fat", "monounsaturated_fat", "polyunsaturated_fat",
		"portion", "potassium", "protein", "saturated_fat", "sodium", "sugar"}

	var values []float64
	for _, field := range fields {
		val, ok := parseFloat(ctx, field)
		if !ok {
			return
		}
		values = append(values, val)
	}

	var existingFood models.Food
	if err := database.DB.Where("name = ?", name).First(&existingFood).Error; err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Name already used."})
		return
	}

	var filename string
	if file, err := ctx.FormFile("image"); err == nil {
		ext := filepath.Ext(file.Filename)
		filename = fmt.Sprintf("%s%s", name, ext)
		savePath := fmt.Sprintf("../web/public/uploads/%s", filename)

		if err := ctx.SaveUploadedFile(file, savePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to upload file."})
			return
		}
	}

	calorie, carbohydrate, cholesterol := values[0], values[1], values[2]
	fat, monounsaturatedFat, polyunsaturatedFat := values[3], values[4], values[5]
	portion, potassium, protein := values[6], values[7], values[8]
	saturatedFat, sodium, sugar := values[9], values[10], values[11]

	food := models.Food{
		Name:               &name,
		Calorie:            &calorie,
		Carbohydrate:       &carbohydrate,
		Category:           &category,
		Cholesterol:        &cholesterol,
		Fat:                &fat,
		Image:              &filename,
		MonounsaturatedFat: &monounsaturatedFat,
		PolyunsaturatedFat: &polyunsaturatedFat,
		Portion:            &portion,
		Potassium:          &potassium,
		Protein:            &protein,
		SaturatedFat:       &saturatedFat,
		Sodium:             &sodium,
		Sugar:              &sugar,
		Unit:               &unit,
	}

	if err := database.DB.Create(&food).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't create data."})
		return
	}

	ctx.JSON(http.StatusOK, food)
}
