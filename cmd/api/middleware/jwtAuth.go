package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/muhammadsaefulr/escommerce/internal/modules/auth"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const SchemaBearer = "Bearer "
		header := ctx.GetHeader("Authorization")

		if header == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Jwt Authorization Token"})
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(header, SchemaBearer) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Authorization Header"})
			ctx.Abort()
			return
		}

		tokenStr := header[len(SchemaBearer):]
		claims := &auth.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return auth.JwtKey, nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Jwt Token"})
			ctx.Abort()
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Jwt Token"})
			ctx.Abort()
			return
		}

		ctx.Set("username", claims.Username)
		ctx.Set("user_id", claims.UserId)
		ctx.Next()
	}
}
