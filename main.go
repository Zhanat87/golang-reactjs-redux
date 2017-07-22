package main

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"time"
	ginjwt "github.com/appleboy/gin-jwt"
	"gopkg.in/dgrijalva/jwt-go.v3"
	_ "github.com/joho/godotenv/autoload"
	"fmt"
)

func main() {
	engine := gin.Default()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// configuration for static files and templates
	engine.LoadHTMLGlob("./templates/*.html")
	engine.StaticFile("/favicon.ico", "./favicon.ico")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Advanced Cloud Native Go with Gin Framework",
		})
	})

	// the jwt middleware
	authMiddleware := &ginjwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		//Timeout:    time.Hour,
		//MaxRefresh: time.Hour,
		Timeout:    time.Minute,
		MaxRefresh: time.Minute,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			if (userId == "admin" && password == "password") || (userId == "test" && password == "test") {
				return userId, true
			}

			return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			if userId == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,

		// roles:
		// https://github.com/appleboy/gin-jwt/blob/f7347d7c16fc8297bf5b7215b9e4847a5ca05d90/auth_jwt_test.go#L442
		PayloadFunc: func(userId string) map[string]interface{} {
			// Set custom claim, to be checked in Authorizator method
			return map[string]interface{}{"user": userId, "role": "ADMIN"}
		},
	}

	engine.POST("/api/login", authMiddleware.LoginHandler)

	engine.Use(authMiddleware.MiddlewareFunc())
	engine.POST("/api/logout", func(c *gin.Context) {
		claims, _ := c.Get("JWT_PAYLOAD")
		userID := claims.(jwt.MapClaims)["id"].(string)
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s logged out", userID)})
	})

	auth := engine.Group("/auth")

	// the hello message endpoint with JSON response from map
	auth.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin Framework."})
	})

	auth.GET("/refresh_token", authMiddleware.RefreshHandler)


	engine.Run(port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
