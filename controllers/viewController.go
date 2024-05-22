package controllers

import (
	"appconf"
	"global"
	"models"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Root (c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func Version(c *gin.Context) {

	info := time.Now().Format("2006-01-02 15:04:05")

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Login Api by LAHB",
		"info1": "Login Api by LAHB",
		"info2": "🇳🇴 " + info,
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

	user_url := global.GetUserUrl(global.IntToString(c.Keys["user"].(models.Users).Id))

	c.HTML(http.StatusOK, "user_home.html", gin.H{
		"title": "Home",
		"user":  c.Keys["user"],
		"css":   "user.css",
		"js":    "user_home.js",
		"url":   user_url,
		"act":   global.ActToString(c.Keys["user"].(models.Users).AccessTime),
	})
}

func ViewNewUsers(c *gin.Context) {
	c.HTML(http.StatusOK, "newusers.html", gin.H{
		"title":    "New users",
		"user":     c.Keys["user"],
		"css":      "user.css",
		"js":       "users.js",
		"ginmode":  appconf.GetVal("gin_mode"),
		"act":      global.ActToString(c.Keys["user"].(models.Users).AccessTime),
		"newusers": global.GetNewUsers(),
		"countnew": global.GetCountNewUsers(),
	})
}

func ViewManageUsers(c *gin.Context) {
	c.HTML(http.StatusOK, "users.html", gin.H{
		"title":       "Manage All Users",
		"user":        c.Keys["user"],
		"css":         "user.css",
		"js":          "users.js",
		"ginmode":     appconf.GetVal("gin_mode"),
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
	user_url := global.GetUserUrl(uid)
	c.HTML(http.StatusOK, "edit_user.html", gin.H{
		"title":   "Edit User",
		"user":    c.Keys["user"],
		"edituid": global.GetUser(uid),
		"css":     "user.css",
		"js":      "edit_user.js",
		"ginmode": appconf.GetVal("gin_mode"),
		"url":     user_url,
		"act":     global.ActToString(global.GetUser(uid).AccessTime),
	})

}

func AppStart(c *gin.Context) {
	c.HTML(http.StatusOK, "appstart.html", gin.H{
		"title": "App Start",
		"user":  c.Keys["user"],
		"css":   "user.css",
		"js":    "appstart.js",
	})
}
