package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func (m *DefaultMiddleware) AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		authorizationString := strings.Split(authorization, " ")

		if len(authorizationString) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if strings.ToLower(authorizationString[0]) != "basic" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		token, _ := jwt.Parse(authorizationString[1], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, nil
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SIGNED_STRING")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["user_id"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

	}

}
