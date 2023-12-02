package middleware

import (
	security "api/security"

	"github.com/gin-gonic/gin"
)

func ValidateJWT(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	token := security.ExtractTokenFromHeader(bearerToken)

	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "UnAuthorized"})
	}

	if err := security.VerifyJWT(token); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unathorized: invalid token"})
	}

	c.Next()
}

func IsAdmin(c *gin.Context) {
	header := c.Request.Header.Get("Authorization")
	token := security.ExtractTokenFromHeader(header)
	claims, err := security.GetClaim(token, "userType")
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": err})
	}

	if claims == "" {
		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
	}

	if claims == "0" {
		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
	}

	c.Next()
}
