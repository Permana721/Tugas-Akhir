package routes

import (
	"api/configs/app_config"
	student_controller "api/controllers/user_controller"
	"api/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
    route := app
    app.Use(app_config.CORSConfig())

    route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)
    route.POST("/student/login", student_controller.Login)

    // Lindungi route dengan middleware AuthMiddleware
    protected := route.Group("/")
    protected.Use(middleware.AuthMiddlewares())
    {
        route.POST("/student/add", student_controller.Store)
        protected.GET("/student", student_controller.GetAllStudent)
        protected.GET("/student/:id", student_controller.GetById)
        protected.PATCH("/student/update/:id", student_controller.Update)
        protected.DELETE("/student/delete/:id", student_controller.Delete)
        protected.POST("/file", student_controller.HandleUploadFile)
    }
}