package global

import (
	"initializers"
	"models"
	"strconv"

	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func IsValidEmail(email string) error {
	var err error

	if email == "" {
		err = errors.New("email must be longer than 0 characters")
		return err
	}

	if len(email) > 100 {
		err = errors.New("email must be less than 100 characters")
		return err
	}

	if len(email) < 9 {
		err = errors.New("email must be valid")
		return err
	}

	if !strings.Contains(email, "@") {
		err = errors.New("email must be valid")
		return err
	}

	if !strings.Contains(email, ".") {
		err = errors.New("email must be valid")
		return err
	}

	if strings.Contains(email, " ") {
		err = errors.New("email must be valid")
		return err
	}

	return nil
}

func IsValidPassword(password string) error {
	var err error

	if password == "" {
		err = errors.New("Password must be longer than 0 characters")
		return err
	}

	if len(password) < 8 {
		err = errors.New("password must be at least 8 characters")
		return err
	}

	if len(password) > 50 {
		err = errors.New("password must be less than 50 characters")
		return err
	}

	return nil
}

func EmailExists(email string) bool {
	var user models.Users
	if err := initializers.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return false
	}
	return true
}

func UintToString(uid uint) string {
	return strconv.Itoa(int(uid))
}

func UserToUInt(uid string) (uint, error) {
	var err error

	if uid == "" {
		err = errors.New("uid must be longer than 0 characters")
		return 0, err
	}

	id, err := strconv.Atoi(uid)
	if err != nil {
		err = errors.New("uid must be a number")
		return 0, err
	}

	return uint(id), nil
}

func GetEmailFromUid(uid uint) string {
	var user models.Users
	initializers.DB.Where("id = ?", uid).First(&user)
	return user.Email
}

func UIntToUser(uid uint) (models.Users, error) {
	var err error
	var user models.Users

	if uid == 0 {
		err = errors.New("uid must be longer than 0 characters")
		return user, err
	}

	initializers.DB.Where("id = ?", uid).First(&user)
	if user.Id == 0 {
		err = errors.New("user not found")
		return user, err
	}

	return user, nil
}

func EmailToUser(email string) (models.Users, error) {
	var err error
	var user models.Users

	if email == "" {
		err = errors.New("email must be longer than 0 characters")
		return user, err
	}
	
	initializers.DB.Where("email = ?", email).First(&user)
	if user.Id == 0 {
		err = errors.New("user not found")
		return user, err
	}

	return user, nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CountUsers() int {
	var count int64
	initializers.DB.Model(&models.Users{}).Count(&count)
	return int(count)
}

func GetUsers() []models.Users {
	var users []models.Users
	initializers.DB.Find(&users)
	return users
}