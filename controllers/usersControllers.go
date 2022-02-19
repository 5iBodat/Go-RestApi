package controllers

import (
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// struct input
type Users struct {
	Id              int       `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirm_password"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func GetUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []Users
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user Users
	db.Find(&user, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses validasi input
	if input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name Must Be Filled"})
		return
	}

	//proses validasi email
	if !isEmailValid(input.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email tidak valid"})
		return
	}

	//proses validasi password
	if !passwordLength(input.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password minimal 8 karakter"})
		return
	}

	if input.Password != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password tidak sama"})
		return
	}

	// proses hash password
	hash, _ := hashPassword(input.Password)

	// proses create user
	user := Users{
		Name:     input.Name,
		Email:    input.Email,
		Password: hash,
	}

	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})

}

func hashPassword(password string) (string, error) {
	//proses hash password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

var emailregex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailregex.MatchString(e)
}

func passwordLength(p string) bool {
	if len(p) < 8 {
		return false
	}
	return true
}
