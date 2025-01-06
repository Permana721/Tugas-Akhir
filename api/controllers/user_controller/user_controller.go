package student_controller

import (
	"time"
	"api/database"
	"api/models"
	"api/request"
	"api/responses"
	"fmt"
	"net/http"
	"os"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	// "gorm.io/gorm"


	"github.com/gin-gonic/gin"
)

func GetAllStudent(ctx *gin.Context) {
	// Mendeklarasikan variabel students yang akan menampung data siswa
	students := new([]models.Student)

	// Menarik data dari tabel students di database
	err := database.DB.Table("students").Find(&students).Error

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
		"data": students, // Mengembalikan data siswa yang berhasil diambil
	})
}


func GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	student := new(responses.StudentResponse)

	errDb := database.DB.Table("students").Where("id = ?", id).Find(&student).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	if student.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message" : "data transmitted",
		"data": student,
	})
}

func Store(ctx *gin.Context) {
	studentReq := new(request.StudentRequest)

	if errReq := ctx.ShouldBind(&studentReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	studentEmailExist := new(models.Student)
	database.DB.Table("students").Where("email = ?", studentReq.Email).First(&studentEmailExist)

	if studentEmailExist.Email != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already used.",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(studentReq.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash password.",
		})
		return
	}

	students := new(models.Student)
	students.Name = &studentReq.Name
	students.Email = &studentReq.Email
	hashedPasswordStr := string(hashedPassword)
	students.Password = &hashedPasswordStr
	students.Age = &studentReq.Age

	errDb := database.DB.Table("students").Create(&students).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't create data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data Successfully Created",
		"data": students,
	})
}

func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	student := new(models.Student)
	studentReq := new(request.StudentRequest)
	studentEmailExist := new(models.Student)

	// Bind data dari request ke struct
	if errReq := ctx.ShouldBind(&studentReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	// Cek apakah data student dengan ID yang diberikan ada
	errDb := database.DB.Table("students").Where("id = ?", id).Find(&student).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if student.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}

	// Cek apakah email sudah digunakan oleh user lain
	errStudentEmailExist := database.DB.Table("students").Where("email = ?", studentReq.Email).Find(&studentEmailExist).Error
	if errStudentEmailExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if studentEmailExist.Email != nil && *student.ID != *studentEmailExist.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already used",
		})
		return
	}

	// Jika password diisi, hash password baru
	if studentReq.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(studentReq.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to hash password",
			})
			return
		}
		hashedPasswordStr := string(hashedPassword)
		student.Password = &hashedPasswordStr
	}

	// Update data student
	student.Name = &studentReq.Name
	student.Email = &studentReq.Email
	student.Age = &studentReq.Age

	errUpdate := database.DB.Table("students").Where("id = ?", id).Updates(&student).Error
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't update data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data updated successfully",
		"data":    student,
	})
}


func Delete(ctx *gin.Context) {
    id := ctx.Param("id")
    student := new(models.Student)

    result := database.DB.Table("students").Where("id = ?", id).Find(&student)
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

    errDb := database.DB.Table("students").Where("id = ?", id).Delete(&models.Student{}).Error
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
	loginReq := new(request.LoginRequest)

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
		})
		return
	}

	user := new(models.Student)
	err := database.DB.Table("students").Where("email = ?", loginReq.Email).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(loginReq.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	// Buat token JWT
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

	// Kirim respon dengan token
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}

func HandleUploadFile(ctx *gin.Context) {
    // Ambil file dari request
    fileHeader, err := ctx.FormFile("file")
    if err != nil {
        ctx.AbortWithStatusJSON(400, gin.H{
            "message": "file is required",
        })
        return
    }

    // Pastikan direktori penyimpanan ada
    os.MkdirAll("./public/files", os.ModePerm)

    // Simpan file ke direktori public/files
    savePath := fmt.Sprintf("./public/files/%s", fileHeader.Filename)
    errUpload := ctx.SaveUploadedFile(fileHeader, savePath)
    if errUpload != nil {
        ctx.JSON(500, gin.H{
            "message": "internal server error, can't save file",
        })
        return
    }

    // Kirimkan respons sukses
    ctx.JSON(200, gin.H{
        "message": "file uploaded",
        "path":    savePath,
    })
}