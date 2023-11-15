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

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	go r.GET("/", controllers.Index)
	
	go r.POST("/signup", controllers.SignUp)
	go r.GET("/signup", controllers.ViewSignup)

	go r.POST("/login", controllers.ViewLogin)
	go r.GET("/login", controllers.ViewLogin)

	userRoutes := r.Group("/u")
	{
		userRoutes.Use(middleware.RequireAuth)
		go userRoutes.POST("/logout", controllers.Logout)
		go userRoutes.GET("/u", controllers.GetUsers)
		go userRoutes.GET("/:id", controllers.GetUser)
		go userRoutes.PUT("/:id", controllers.UpdateUser)
		go userRoutes.DELETE("/:id", controllers.DeleteUser)

	}

	viewRoutes := r.Group("/v")
	{
		viewRoutes.Use(middleware.RequireAuth)
		go viewRoutes.GET("/home", controllers.ViewUserHome)
		go viewRoutes.GET("/users", controllers.ViewManageUsers)
	}

	return r
}
