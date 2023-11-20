package controllers

import (
	"filefunc"
	"fmt"
	"initializers"
	"global"

	"log"
	"models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email     string `json:"email"`
		Password string `json:"password"`
		Password2 string `json:"password2"`
	}
	log.Print(body)
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body"})
		return
	}

	// validate the email/password
	email := body.Email
	password := body.Password
	password2 := body.Password2

	var err error
	err = global.IsValidEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		return
	}

	if password != password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Passwords do not match"})
		return
	}

	err = global.IsValidPassword(password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		return
	}
	
	// check if the email is already in use
	if global.EmailExists(email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email is already in use"})
		return
	}

	// hash the password before save
	hashedPassword, err := global.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash password"})
		return
	}

	// first user gets to be admin
	role := "user"
	isauth := false
	if global.CountUsers() == 0 {
		role = "admin"
		isauth = true
	}

	// create a user
	user := models.Users{
		Email:    email,
		Password: string(hashedPassword),
		Role: role,
		IsAuth: isauth,
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create user"})
		return
	}

	// return the user
	c.JSON(http.StatusOK, gin.H{
		"message": "success"})

}

func Login(c *gin.Context) {

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		IsAuth bool `json:"isauth"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body"})
		return
	}

	// validate the email/password
	email := body.Email
	password := body.Password

	var err error
	err = global.IsValidEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		return
	}

	err = global.IsValidPassword(password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		return
	}

	// find the user by email
	user, err := global.EmailToUser(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Email or password is invalid"})
		return
	}
	
	if !global.EmailExists(email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email or password is invalid"})
		return
	}

	if !global.CheckPasswordHash(password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email or password is invalid"})
		return
	}

	if !user.IsAuth {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User is not authenticated"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to generate token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("goAuth", tokenString, 3600*24*30, "/", "", false, true)

	usrFileFolder := fmt.Sprintf(os.Getenv("SHARE_FOLDER") + "/%d/files", user.Id)
	
	if !filefunc.IsExists(usrFileFolder) {
		filefunc.CreateFolder(usrFileFolder)
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("goAuth", "", 0, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func GetUsers(c *gin.Context) {
	var users []models.Users
	if err := initializers.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.Users
	if err := initializers.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// func UpdateUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var user models.Users
// 	if err := initializers.DB.Where("id = ?", id).First(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "failed to get user"})
// 		return
// 	}

// 	var body struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	if c.BindJSON(&body) != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "failed to read body"})
// 		return
// 	}

// 	if body.Email != "" {
// 		user.Email = body.Email
// 	}

// 	if body.Password != "" {
// 		user.Password = body.Password
// 	}

// 	if err := initializers.DB.Save(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "failed to update user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }

func DeleteUser(c *gin.Context) {
	var body struct {
		Id    string `json:"id"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body"})
		return
	}

	uid := body.Id

	if uid == "1" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Not allowed to delete superadmin!"})
		return
	}

	var user models.Users
	if err := initializers.DB.Where("id = ?", uid).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get user"})
		return
	}

	if err := initializers.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, user)
}

func UpdateAuth(c *gin.Context) {
	var body struct {
		Id    string `json:"id"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body"})
		return
	}

	var user models.Users
	if err := initializers.DB.Where("id = ?", body.Id).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get user"})
		return
	}

	// protect superadmin
	if body.Id == "1" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can´t auth for change superadmin!"})
		return
	}

	switch user.IsAuth {
	case true:
		user.IsAuth = false
	case false:
		user.IsAuth = true
	}

	if err := initializers.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update auth"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success"})
	
}

func UpdateRole(c *gin.Context) {
	var body struct {
		Id    string `json:"id"`
		Role string `json:"role"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body"})
		return
	}
	
	var user models.Users
	if err := initializers.DB.Where("id = ?", body.Id).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get user"})
		return
	}

	if body.Id == "1" {
		http.Redirect(c.Writer, c.Request, "/v/users", http.StatusMovedPermanently)
		return
	}

	user.Role = body.Role
	// log.Println(user.Role)

	if err := initializers.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success"})
	
}

func SetNewPassword(c *gin.Context) {
	var body struct {
		Id    string `json:"id"`
		Password string `json:"password"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body"})
		return
	}

	var user models.Users
	if err := initializers.DB.Where("id = ?", body.Id).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get user"})
		return
	}

	if global.IsValidPassword(body.Password) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Password is not valid"})
		return
	}

	hashedPassword, err := global.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash password"})
		return
	}

	user.Password = string(hashedPassword)
	if err := initializers.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success"})

}