package user_controller

import (
	"api/database"
	models "api/models/user_model"
	"api/responses"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"


	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {
	users := new([]models.User)
	err := database.DB.Table("users").Find(&users).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, users)
}


func GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(responses.UserResponse)

	errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}

	ctx.JSON(200, user)
}

func Store(ctx *gin.Context) {
    name := ctx.PostForm("name")
    email := ctx.PostForm("email")
    password := ctx.PostForm("password")
    ageStr := ctx.PostForm("age")

    if name == "" || email == "" || password == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "message": "Semua field harus diisi!",
        })
        return
    }

    age, err := strconv.Atoi(ageStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "message": "Age harus berupa angka!",
        })
        return
    }

    userEmailExist := new(models.User)
    database.DB.Table("users").Where("email = ?", email).First(&userEmailExist)

    if userEmailExist.Email != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "message": "Email already used.",
        })
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "message": "Failed to hash password.",
        })
        return
    }

    hashedPasswordStr := string(hashedPassword)
    user := models.User{
        Name:     &name,
        Email:    &email,
        Password: &hashedPasswordStr,
        Age:      &age,
    }

    errDb := database.DB.Table("users").Create(&user).Error
    if errDb != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "message": "Can't create data.",
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Data Successfully Created",
        "data":    user,
    })
}

func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)
	userEmailExist := new(models.User)

	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	ageStr := ctx.PostForm("age")

	if name == "" || email == "" || password == "" || ageStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Semua field harus diisi!",
		})
		return
	}

	age, err := strconv.Atoi(ageStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "message": "Age harus berupa angka!",
        })
        return
    }

	errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}

	erruserEmailExist := database.DB.Table("users").Where("email = ?", email).Find(&userEmailExist).Error
	if erruserEmailExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if userEmailExist.Email != nil && *user.ID != *userEmailExist.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already used",
		})
		return
	}

	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to hash password",
			})
			return
		}
		hashedPasswordStr := string(hashedPassword)
		user.Password = &hashedPasswordStr
	}

	user.Name = &name
	user.Email = &email
	user.Age = &age

	errUpdate := database.DB.Table("users").Where("id = ?", id).Updates(&user).Error
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't update data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data updated successfully",
		"data":    user,
	})
}

func Delete(ctx *gin.Context) {
    id := ctx.Param("id")
    user := new(models.User)

    result := database.DB.Table("users").Where("id = ?", id).Find(&user)
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

    errDb := database.DB.Table("users").Where("id = ?", id).Delete(&models.User{}).Error
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

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	if email == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Semua field harus diisi!",
		})
		return
	}

	user := new(models.User)
	err := database.DB.Table("users").Where("email = ?", email).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}

func HandleUploadFile(ctx *gin.Context) {
    fileHeader, err := ctx.FormFile("file")
    if err != nil {
        ctx.AbortWithStatusJSON(400, gin.H{
            "message": "file is required",
        })
        return
    }

    os.MkdirAll("./public/files", os.ModePerm)

    savePath := fmt.Sprintf("./public/files/%s", fileHeader.Filename)
    errUpload := ctx.SaveUploadedFile(fileHeader, savePath)
    if errUpload != nil {
        ctx.JSON(500, gin.H{
            "message": "internal server error, can't save file",
        })
        return
    }

    ctx.JSON(200, gin.H{
        "message": "file uploaded",
        "path":    savePath,
    })
}