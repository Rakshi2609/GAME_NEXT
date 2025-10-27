package middleware

import (
	"net/http"
	"os"
	"strings"

	"codechef/backend/utils"

	"github.com/gin-gonic/gin"
)

type ContextKey string

const (
	CtxUserID ContextKey = "uid"
	CtxRole   ContextKey = "role"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")
		if !strings.HasPrefix(strings.ToLower(bearer), "bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			return
		}
		token := strings.TrimSpace(bearer[len("Bearer "):])
		claims, err := utils.ParseJWT(os.Getenv("JWT_SECRET"), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.Set(string(CtxUserID), claims.UserID)
		c.Set(string(CtxRole), claims.Role)
		c.Next()
	}
}

func RoleAllowed(roles ...string) gin.HandlerFunc {
	allowed := map[string]struct{}{}
	for _, r := range roles {
		allowed[r] = struct{}{}
	}
	return func(c *gin.Context) {
		role, _ := c.Get(string(CtxRole))
		if _, ok := allowed[role.(string)]; !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}
		c.Next()
	}
}
