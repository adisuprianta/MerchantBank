package controllers

import (
	"MerchantBank/app"
	"MerchantBank/models"
	"fmt"
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

	logoModel := models.ActivityLog{Action: "Create", UserId: user.ID, Description: "Membuat akun baru"}

	app.DB.Create(&logoModel)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"massage": "Successfully create user",
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

	logoModel := models.ActivityLog{Action: "Login", UserId: user.ID, Description: "Melakukan Login"}

	app.DB.Create(&logoModel)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"massage": "successfully login ",
		"token":   tokenString,
	})
}

func Logout(c *gin.Context) {

	data, exits := c.Get("user")

	if exits {
		if user, ok := data.(models.User); ok {
			app.DB.First(&user, "email = ?", user.Email)

			if user.ID == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			fmt.Println(user)

			err := app.DB.Model(&user).Updates(map[string]interface{}{
				"email":    user.Email,
				"password": user.Password,
				"enabled":  false,
			}).Error
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Failed to create token",
				})
				return
			}

			logoModel := models.ActivityLog{Action: "Logout", UserId: user.ID, Description: "Melakukan logout"}

			app.DB.Create(&logoModel)
			c.JSON(http.StatusCreated, gin.H{
				"status":  http.StatusOK,
				"massage": `successfully logout `,
			})
			return
		}
	}

	c.AbortWithStatus(http.StatusUnauthorized)

}
