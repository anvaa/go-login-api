package routers

import (
	"controllers"
	"middleware"
	
	"os"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {

	if os.Getenv("GIN_MODE") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	if os.Getenv("GIN_MODE") == "debug" {
		r.Static("/css", "embedfiles/web/css")
		r.Static("/js", "embedfiles/web/js")
		r.Static("/share", "embedfiles/share")
		r.LoadHTMLGlob("embedfiles/web/html/*.html")
	} else {
		r.Static("/css", ".static/css")
		r.Static("/js", ".static/js")
		r.Static("/share", ".static/share")
		r.LoadHTMLGlob(".static/html/*.html")	
	}

	r.GET("/", controllers.Index)

	r.POST("/signup", controllers.ViewSignup)
	r.GET("/signup", controllers.ViewSignup)

	r.POST("/login", controllers.ViewLogin)
	r.GET("/login", controllers.ViewLogin)

	r.POST("/logout", controllers.Logout)

	userRoutes := r.Group("/user")
	{
		userRoutes.Use(middleware.RequireAuth)
		userRoutes.Use(middleware.IsAdmin)
		userRoutes.Use(middleware.IsAuth)

		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUser)
		// go userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.POST("/delete/:id", controllers.DeleteUser)
		userRoutes.POST("/auth", controllers.UpdateAuth)
		userRoutes.POST("/role", controllers.UpdateRole)
		userRoutes.POST("/psw", controllers.SetNewPassword)
		userRoutes.POST("/act", controllers.SetAct)
	}

	viewRoutes := r.Group("/v")
	{
		viewRoutes.Use(middleware.RequireAuth)
		viewRoutes.Use(middleware.IsAuth)

		// not admin
		viewRoutes.GET("/home", controllers.ViewHome)
		viewRoutes.GET("/userhome", controllers.ViewUserHome)

		// is admin
		viewRoutes.GET("/adminhome", middleware.IsAdmin, controllers.ViewAdminHome)
		viewRoutes.GET("/users", middleware.IsAdmin, controllers.ViewManageUsers)
		viewRoutes.GET("/user/:id", middleware.IsAdmin, controllers.ViewEditUser)
	}

	return r
}
