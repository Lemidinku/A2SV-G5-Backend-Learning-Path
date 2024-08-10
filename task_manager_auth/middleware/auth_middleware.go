package middleware

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	  // TODO: Implement JWT validation logic
	  authHeader := c.GetHeader("Authorization")
	  if authHeader == "" {
		c.JSON(401, gin.H{"error": "Authorization header is required"})
		c.Abort()
		return
	  }
	  var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	  authParts := strings.Split(authHeader, " ")
	  if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		c.JSON(401, gin.H{"error": "Invalid authorization header"})
		c.Abort()
		return
	  }
	  
	  token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		  return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
	  
		return jwtSecret, nil
	  })
	  
	  if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Invalid JWT"})
		c.Abort()
		return
	  }
	  c.Next()
	}
  }


  func OnlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
		authParts := strings.Split(authHeader, " ")
		
		token, _:= jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
		
			return jwtSecret, nil
		})
		
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid JWT claims"})
			c.Abort()
			return
		}

		role, ok := claims["Role"].(string)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid JWT role"})
			c.Abort()
			return
		}

		if role != "admin" {
			c.JSON(403, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
	}
  }

