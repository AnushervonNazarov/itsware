package controllers

import (
	"context"
	"itsware/db"
	"itsware/internal/constants"
	"itsware/internal/services"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "userID"
	userRoleCtx         = "userRole"
)

// Middleware configuration
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
			return
		}

		// Parse token with custom claims
		token, err := jwt.ParseWithClaims(tokenString, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims, ok := token.Claims.(*services.CustomClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}

		c.Set("claims", claims) // Store *CustomClaims
		c.Next()
	}
}

// SetDBSessionVariables.go
func SetDBSessionVariables() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Retrieve claims from context
		claimsInterface, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		claims, ok := claimsInterface.(*services.CustomClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid claims type"})
			return
		}

		// 2. Acquire database connection
		conn, err := db.Pool.Acquire(c.Request.Context())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}

		// 3. Set session variables
		_, err = conn.Exec(c.Request.Context(), `
			SELECT 
				set_config('myapp.session.user_id', $1, false),
				set_config('myapp.session.tenant_id', $2, false),
				set_config('myapp.session.user_role', $3, false)`,
			strconv.Itoa(claims.UserID),
			strconv.Itoa(claims.TenantID),
			claims.Role,
		)
		if err != nil {
			conn.Release()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to set session variables"})
			return
		}

		// 4. Store connection in request context
		ctx := context.WithValue(c.Request.Context(), constants.DBConnKey, conn)
		c.Request = c.Request.WithContext(ctx)

		// 5. Release connection AFTER request completes
		defer func() {
			conn.Release()
		}()

		c.Next()
	}
}
