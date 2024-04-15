package routers

import (
	"appconf"
	"controllers"
	"middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {

	if appconf.GetVal("gin_mode") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	if appconf.GetVal("gin_mode") == "debug" {
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
		
		userRoutes.POST("/delete/:id", controllers.DeleteUser)
		userRoutes.POST("/auth", controllers.UpdateAuth)
		userRoutes.POST("/role", controllers.UpdateRole)
		userRoutes.POST("/psw", controllers.SetNewPassword)
		userRoutes.POST("/act", controllers.SetAct)
		userRoutes.POST("/url", controllers.UpdateUrl)
	}

	viewRoutes := r.Group("/v")
	{
		viewRoutes.Use(middleware.RequireAuth)
		viewRoutes.Use(middleware.IsAuth)

		// not admin
		viewRoutes.GET("/appstart", controllers.AppStart)

		// is admin
		viewRoutes.GET("/newusers", middleware.IsAdmin, controllers.ViewNewUsers)
		viewRoutes.GET("/users", middleware.IsAdmin, controllers.ViewManageUsers)
		viewRoutes.GET("/user/:id", middleware.IsAdmin, controllers.ViewEditUser)
		viewRoutes.GET("/userhome", middleware.IsAdmin, controllers.ViewNewUsers)
	}

	return r
}
