package middleware

import (
	"MerchantBank/app"
	"MerchantBank/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func RequireAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[7:]
	}
	//fmt.Println(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.User

		app.DB.First(&user, claims["sub"])

		if user.ID == 0 || !user.Enabled {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()
	} else {
		fmt.Println(err)
	}

}
