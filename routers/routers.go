package routers

import (
	"controllers"
	"middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(wd string) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies([]string{""})

	r.Static("/css", wd + "/templates/css")
	r.Static("/js", wd + "/templates/js")
	r.Static("/share", wd + "/share")
	r.LoadHTMLGlob(wd + "/templates/*.gohtml")


	go r.GET("/", controllers.Index)
	
	go r.POST("/signup", controllers.ViewSignup)
	go r.GET("/signup", controllers.ViewSignup)

	go r.POST("/login", controllers.ViewLogin)
	go r.GET("/login", controllers.ViewLogin)

	go r.GET("/logout", controllers.Logout)

	userRoutes := r.Group("/user")
	{
		userRoutes.Use(middleware.RequireAuth)
		userRoutes.Use(middleware.IsAdmin)
		
		go userRoutes.GET("/", controllers.GetUsers)
		go userRoutes.GET("/:id", controllers.GetUser)
		go userRoutes.PUT("/:id", controllers.UpdateUser)
		go userRoutes.POST("/delete/:id", controllers.DeleteUser)
		go userRoutes.POST("/auth/:id", controllers.UpdateAuth)
		go userRoutes.POST("/role", controllers.UpdateRole)
		go userRoutes.POST("/psw", controllers.SetNewPassword)

	}

	viewRoutes := r.Group("/v")
	{
		viewRoutes.Use(middleware.RequireAuth)
		go viewRoutes.GET("/home", controllers.ViewUserHome)
		go viewRoutes.GET("/users", middleware.IsAdmin, controllers.ViewManageUsers)
		go viewRoutes.GET("/user/:id", middleware.IsAdmin, controllers.ViewEditUser)
	}

	return r
}
