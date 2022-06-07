// routes.go

package routes

import (
	"log"

	"github.com/a11dev/go-gen-backend/internal/config"
	"github.com/a11dev/go-gen-backend/internal/handlers"
	"github.com/a11dev/go-gen-backend/internal/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/store"
)

func InitializeRoutes(router *gin.Engine, cfg *config.Config, authenticator auth.Authenticator, store store.Cache) {
	// the jwt middleware
	authMiddleware, err := middleware.JwtAuth(authenticator, cfg)
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	router.POST("/login", authMiddleware.LoginHandler)
	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := router.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/healthcheck", handlers.HealthCheck)
		auth.GET("/task/:id", handlers.GetTask)
		auth.GET("/tasks/:filterid", handlers.GetTasks)
	}
}
