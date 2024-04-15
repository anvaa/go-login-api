package global

import (
	"fmt"
	"initializers"
	"models"

	"errors"
	"regexp"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func IsValidEmail(email string) error {
	var err error

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		err = errors.New("not a valid username")
		return err
	}

	return nil
}

func IsValidPassword(password string) error {
	var err error

	if password == "" {
		err = errors.New("password must be longer than 0 characters")
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

func IntToString(uid int) string {
	return strconv.Itoa(uid)
}

func GetEmailFromUid(uid int) string {
	var user models.Users
	initializers.DB.Where("id = ?", uid).First(&user)
	return user.Email
}

func GetUserUrl(uid string) string {
	var url models.Links
	initializers.DB.Where("user_id = ?", uid).First(&url)

	if url.Url == "" {
		url.Url = "/login"
		initializers.DB.Create(&models.Links{UserId: StringToInt(uid), Url: url.Url})
		return url.Url
	}
	
	return url.Url
}

func EmailToUser(email string) (models.Users, error) {
	var err error
	var user models.Users

	if email == "" {
		err = errors.New("username must be longer than 0 characters")
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

func GetAuthUsers() []models.Users {
	var users []models.Users
	initializers.DB.Where("is_auth = ?", true).Find(&users)
	return users
}

func GetCountAuthUsers() int {
	var count int64
	initializers.DB.Model(&models.Users{}).Where("is_auth = ?", true).Count(&count)
	return int(count)
}

func GetUnauthUsers() []models.Users {
	var users []models.Users
	initializers.DB.Where("is_auth = ?", false).Find(&users)
	return users
}

func GetCountUnauthUsers() int {
	var count int64
	initializers.DB.Model(&models.Users{}).Where("is_auth = ?", false).Count(&count)
	return int(count)
}

func GetNewUsers() []models.Users {
	var users []models.Users
	initializers.DB.Where("created_at = updated_at and is_auth = false").Find(&users)
	return users
}

func GetCountNewUsers() int {
	var count int64
	initializers.DB.Model(&models.Users{}).Where("created_at = updated_at and is_auth = false").Count(&count)
	return int(count)
}

func GetUser(uid string) models.Users {
	var user models.Users
	initializers.DB.Where("id = ?", uid).First(&user)
	return user
}

func GetDeletedUsers() []models.Users {
	var users []models.Users
	initializers.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&users)
	return users
}

func GetCountDeletedUsers() int {
	var count int64
	initializers.DB.Model(&models.Users{}).Unscoped().Where("deleted_at IS NOT NULL").Count(&count)
	return int(count)
}

func StringToInt(str string) int {

	if str == "" {
		return 0
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

func ActToString(t int) string {
	
	if t == 0 {
		t = 3600 // 1 hours
	}

	// sec := t % 60
	min := t / 60
	hour := min / 60
	min = min % 60
	day := hour / 24
	hour = hour % 24

	var timeString string
	switch {
	case day > 0:
		timeString = fmt.Sprintf("%d days", day)
		if hour > 0 {
			timeString = fmt.Sprintf("%d days, %d hours", day, hour)
		}
		if min > 0 {
			timeString = fmt.Sprintf("%d days, %d hours, %d minutes", day, hour, min)
		}
	case hour > 0:
		timeString = fmt.Sprintf("%d hours", hour)
		if min > 0 {
			timeString = fmt.Sprintf("%d hours, %d minutes", hour, min)
		}
	case min > 0:
		timeString = fmt.Sprintf("%d minutes", min)
	default:
		timeString = fmt.Sprintf("%d minutes", min)
	}

	return timeString
}

func CalculateAccessTime(t string) int {

	min := StringToInt(t)
	if min == 0 {
		min = 1
	}

	return 60 * min
}
