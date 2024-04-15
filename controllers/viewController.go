package controllers

import (
	"global"
	"models"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func Index(c *gin.Context) {

	info2 := time.Now().Format("2006-01-02 15:04:05") 

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Login Api by LAHB",
		"info1":   "Login Api by LAHB",
		"info2":   "🇳🇴 "+info2,
		"css":   "index.css",
	})
}

func ViewSignup(c *gin.Context) {

	if c.Request.Method == "POST" {
		SignUp(c)
		return
	}

	c.HTML(http.StatusOK, "signup.html", gin.H{
		"title": "Signup",
		"css":   "index.css",
		"js":    "signup.js",
	})
}

func ViewLogin(c *gin.Context) {

	if c.Request.Method == "POST" {
		Login(c)
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
		"user":  c.Keys["user"],
		"css":   "index.css",
		"js":    "login.js",
	})
}

func ViewUserHome(c *gin.Context) {
	c.HTML(http.StatusOK, "user_home.html", gin.H{
		"title": "Home",
		"user":  c.Keys["user"],
		"css":   "user.css",
		"js":    "user_home.js",
		"act":   global.ActToString(c.Keys["user"].(models.Users).AccessTime),
	})
}

func ViewNewUsers(c *gin.Context) {
	c.HTML(http.StatusOK, "newusers.html", gin.H{
		"title":     "New users",
		"user":      c.Keys["user"],
		"css":       "user.css",
		"js":        "users.js",
		"act":       global.ActToString(c.Keys["user"].(models.Users).AccessTime),
		"newusers":  global.GetNewUsers(),
		"countnew":  global.GetCountNewUsers(),
	})
}

func ViewManageUsers(c *gin.Context) {
	c.HTML(http.StatusOK, "users.html", gin.H{
		"title":       "Manage All Users",
		"user":        c.Keys["user"],
		"css":         "user.css",
		"js":          "users.js",
		"authusers":   global.GetAuthUsers(),
		"countauth":   global.GetCountAuthUsers(),
		"unauthusers": global.GetUnauthUsers(),
		"countunauth": global.GetCountUnauthUsers(),
		"delusers":    global.GetDeletedUsers(),
		"countdel":    global.GetCountDeletedUsers(),
	})
}

func ViewEditUser(c *gin.Context) {
	uid := c.Param("id")
	c.HTML(http.StatusOK, "edit_user.html", gin.H{
		"title":   "Edit User",
		"user":    c.Keys["user"],
		"edituid": global.GetUser(uid),
		"css":     "user.css",
		"js":      "edit_user.js",
		"act":     global.ActToString(global.GetUser(uid).AccessTime),
	})

}
