package routes

import (
	"api/configs/app_config"
	user_controller "api/controllers/user_controller"
    blog_controller "api/controllers/blog_controller"
    food_controller "api/controllers/food_controller"
	"api/middleware"

	"github.com/gin-gonic/gin"
	// "golang.org/x/tools/blog"
)

func InitRoute(app *gin.Engine) {
    route := app
    app.Use(app_config.CORSConfig())

    route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)
    route.POST("/user/login", user_controller.Login)

    route.GET("/blog", blog_controller.GetAllBlog)
    route.GET("/blog/:id", blog_controller.GetById)
    route.POST("/blog", blog_controller.Store)
    route.PATCH("/blog/update/:id", blog_controller.Update)
    route.DELETE("/blog/delete/:id", blog_controller.Delete)

    route.GET("/food", food_controller.GetAllFood)
    route.GET("/food/:id", food_controller.GetById)
    route.POST("/food", food_controller.Store)

    // Lindungi route dengan middleware AuthMiddleware
    protected := route.Group("/")
    protected.Use(middleware.AuthMiddlewares())
    {
        route.POST("/user/add", user_controller.Store)
        protected.GET("/user", user_controller.GetAllUser)
        protected.GET("/user/:id", user_controller.GetById)
        route.PATCH("/user/update/:id", user_controller.Update)
        protected.DELETE("/user/delete/:id", user_controller.Delete)
        protected.POST("/file", user_controller.HandleUploadFile)
    }
}