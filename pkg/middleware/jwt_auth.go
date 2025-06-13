package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"simple-go/internal/models"
	_ "simple-go/internal/models"
	"time"
)

// 定义 JWT payload 结构
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("jerry") // 密钥应从配置中读取并妥善保护

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			ajax := models.NewAjaxResult(http.StatusUnauthorized, "请求未授权")
			c.AbortWithStatusJSON(http.StatusUnauthorized, ajax)
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			ajax := models.NewAjaxResult(http.StatusUnauthorized, "Invalid token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, ajax)
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
