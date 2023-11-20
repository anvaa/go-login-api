package routers

import (
	"controllers"
	"middleware"

	"os"
	"github.com/gin-gonic/gin"
)

func SetupRouter(wd string) *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.Default()
	r.SetTrustedProxies(nil)
	// r.TrustedPlatform = "linux"

	r.Static("/css", wd + "/templates/css")
	r.Static("/js", wd + "/templates/js")
	r.Static("/share", wd + "/share")
	r.LoadHTMLGlob(wd + "/templates/*.gohtml")


	go r.GET("/", controllers.Index)
	
	go r.POST("/signup", controllers.ViewSignup)
	go r.GET("/signup", controllers.ViewSignup)

	go r.POST("/login", controllers.ViewLogin)
	go r.GET("/login", controllers.ViewLogin)

	go r.POST("/logout", controllers.Logout)

	userRoutes := r.Group("/user")
	{	
		userRoutes.Use(middleware.RequireAuth)
		userRoutes.Use(middleware.IsAdmin)

		go userRoutes.GET("/", controllers.GetUsers)
		go userRoutes.GET("/:id", controllers.GetUser)
		// go userRoutes.PUT("/:id", controllers.UpdateUser)
		go userRoutes.POST("/delete/:id", controllers.DeleteUser)
		go userRoutes.POST("/auth", controllers.UpdateAuth)
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
