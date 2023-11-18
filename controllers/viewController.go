package controllers

import (
	"global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", gin.H{
		"title": "Home",
		"css":  "index.css",
		"js":   "index.js",
		"logo": "logo.png",
	})
}

func ViewSignup(c *gin.Context) {
	
	if c.Request.Method == "POST" {
		SignUp(c)
		return
	}

	c.HTML(http.StatusOK, "signup.gohtml", gin.H{
		"title": "Signup",
		"css": "user.css",
		"js": "signup.js",
	})
}

func ViewLogin(c *gin.Context) {

	if c.Request.Method == "POST" {
		Login(c)
		return
	}

	c.HTML(http.StatusOK, "login.gohtml", gin.H{
		"title": "Login",
		"user": c.Keys["user"],
		"css": "user.css",
		"js": "login.js",
	})
}

func ViewUserHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.gohtml", gin.H{
		"title": "Home",
		"user": c.Keys["user"],
		"css": "user.css",
		"js": "home.js",
	})
}

func ViewManageUsers(c *gin.Context) {
	c.HTML(http.StatusOK, "auth_users.gohtml", gin.H{
		"title": "Manage Users",
		"user": c.Keys["user"],
		"css": "user.css",
		"js": "auth_users.js",
		"users": global.GetUsers(),
	})
}

func ViewEditUser(c *gin.Context) {
	uid := c.Param("id")
	c.HTML(http.StatusOK, "edit_user.gohtml", gin.H{
		"title": "Edit User",
		"user": c.Keys["user"],
		"edituid": global.GetUser(uid),
		"css": "user.css",
		"js": "edit_users.js",
	})
}