package middleware

import (
	"BrunoyamLesson6/internal/server/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware(signer auth.HS256Signer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid auth header"})
			ctx.Abort() // чтобы дальше на след функцию хендлера не ушло
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := signer.ParseAccessToken(tokenString, auth.ParseOptions{
			ExpectedIssuer:   signer.Issuer,
			ExpectedAudience: signer.Audience,
			AllowMethods:     []string{"HS256"},
			Leeway:           60 * time.Second,
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
