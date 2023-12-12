package controllers

import (
	"MerchantBank/app"
	"MerchantBank/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"net/http"
	"os"
	"time"
)

var authRequest struct {
	Email    string
	Password string
}

func Sigup(c *gin.Context) {

	if c.Bind(&authRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(authRequest.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	user := models.User{Email: authRequest.Email, Password: string(hash), Enabled: false}
	result := app.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"massege": "Successfully create user",
		"email":   authRequest.Email,
	})
}
func Login(c *gin.Context) {
	if c.Bind(&authRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user models.User

	app.DB.First(&user, "email = ?", authRequest.Email)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Failed credential",
		})
		return
	}

	app.DB.Model(&user).Updates(models.User{
		Email:    user.Email,
		Password: user.Password,
		Enabled:  true,
	})

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authRequest.Password))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Failed credential",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  os.Getenv("APP_NAME"),
		"sub":  user.ID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"iat":  time.Now().Unix(),
		"role": "test",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

//
//func Logout(c *gin.Context) {
//	var user models.User
//	user, _ := c.Get("user")
//
//	app.DB.First(&user, "email = ?", user.Email)
//
//	if user.ID == 0 {
//		c.JSON(http.StatusNotFound, gin.H{
//			"error": "Failed credential",
//		})
//		return
//	}
//
//	app.DB.Model(&user).Updates(models.User{
//		Email:    user.Email,
//		Password: user.Password,
//		Enabled:  true,
//	})
//}
