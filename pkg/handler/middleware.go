package handler

import (
	"errors"
	"fmt"
	"mrs_project/pkg/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	roleCtx             = "role"
)

func (h *Handler) AuthMiddleWare(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			c.JSON(401, gin.H{"error": "empty auth header"})
			c.Abort()
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "invalid auth header"})
			c.Abort()
			return
		}

		if len(headerParts[1]) == 0 {
			c.JSON(401, gin.H{"error": "auth token is empty"})
			c.Abort()
			return
		}

		tokenString := headerParts[1]
		claims, err := utils.VerifyAccessToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid or expider access token" + err.Error()})
			c.Abort()
			return
		}

		if requiredRole != "" && claims.Role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			c.Abort()
			return
		}

		c.Set(userCtx, claims.UserID)
		c.Set(roleCtx, claims.Role)

		c.Next()
		fmt.Printf("Request by userID=%d with role=%s processed in %s\n", claims.UserID, claims.Role, time.Since(startTime))
	}
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

func getRole(c *gin.Context) (string, error) {
	role, ok := c.Get(roleCtx)
	if !ok {
		return "", errors.New("user role not found")
	}

	roleString, ok := role.(string)
	if !ok {
		return "", errors.New("user role is of invalid type")
	}

	return roleString, nil
}
