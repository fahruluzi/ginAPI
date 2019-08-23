package middleware

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT_SECRET is string on server for access JWT Token
var JWT_SECRET = "SUPER_SECRET"

// IsAuth adalah fungsi untuk check apakah user sudah authentikasi apa belum
func IsAuth() gin.HandlerFunc {
	return checkJWT()
}

func checkJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil Data dari header Authorization
		authHeader := c.Request.Header.Get("Authorization")
		// explode Token "Bearer 'token....'"
		bearerToken := strings.Split(authHeader, " ")

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// JWT_SECRET is string on server for access JWT Token
			return []byte(JWT_SECRET), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// key yang diprint, hanya diambil dari yang kita kirim saat pembuatan Token
			fmt.Println("hi")
			fmt.Println(claims["user_id"], claims["user_role"])
		} else {
			fmt.Println("hi1")
			fmt.Println(err)
		}
	}
}
