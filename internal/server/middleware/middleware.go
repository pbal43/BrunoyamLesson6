package middleware

import (
	"BrunoyamLesson6/internal/domain"
	"BrunoyamLesson6/internal/server/auth"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
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
			Leeway:           domain.LeewayTimeout,
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

func ZeroLogMiddleware(log *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		duration := time.Since(start)

		log.Info().
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Int("status", ctx.Writer.Status()).
			Str("ip", ctx.ClientIP()).
			Dur("duration", duration).
			Send()
	}
}
