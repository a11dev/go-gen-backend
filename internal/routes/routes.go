// routes.go

package routes

import (
	"log"
	"time"

	"github.com/a11dev/go-gen-backend/internal/config"
	"github.com/a11dev/go-gen-backend/internal/handlers"
	"github.com/a11dev/go-gen-backend/internal/middleware"
	"github.com/a11dev/go-gen-backend/internal/runtimes"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/store"
)

func InitializeRoutes(router *gin.Engine, cfg *config.Config, authenticator auth.Authenticator, store store.Cache, inc chan runtimes.BackendMsg) {
	// the jwt middleware
	authMiddleware, err := middleware.JwtAuth(authenticator, cfg)

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"http://localhost:3000/"}
	// corsConfig.AllowCredentials = true
	// corsConfig.AddAllowHeaders( "access-control-allow-origin"),
	// corsConfig.AllowAllOrigins = true
	// router.Use(cors.New(corsConfig))
	// router.Use(CORS())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		ExposeHeaders:    []string{"Content-Length"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "access-control-allow-origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.POST("/login", authMiddleware.LoginHandler)
	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	chans := router.Group("/chans")
	chans.Use(middleware.Backend1ChannelsMiddleware(inc))
	{
		chans.GET("/routine/:id", handlers.InvokeBackend)
	}

	auth := router.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/healthcheck", handlers.HealthCheck)
		auth.GET("/task/:id", handlers.GetTask)
		auth.GET("/tasks", handlers.GetTasks)
	}
}

// func CORS() gin.HandlerFunc {
// 	// TO allow CORS
// 	return func(c *gin.Context) {
// 		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		// c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
// 		// if c.Request.Method == "OPTIONS" {
// 		// 	c.AbortWithStatus(204)
// 		// 	return
// 		// }
// 		// c.Next()
// 		method := c.Request.Method               //Request method
// 		origin := c.Request.Header.Get("Origin") //Request header
// 		var headerKeys []string                  //Declare request header keys
// 		for k, _ := range c.Request.Header {
// 			headerKeys = append(headerKeys, k)
// 		}
// 		headerStr := strings.Join(headerKeys, ", ")
// 		if headerStr != "" {
// 			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
// 		} else {
// 			headerStr = "access-control-allow-origin, access-control-allow-headers"
// 		}
// 		if origin != "" {
// 			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 			c.Header("Access-Control-Allow-Origin", "*")                                       //This is to allow access to all domains
// 			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //All cross-domain request methods supported by the server, in order to avoid multiple times of browsing Check request
// 			//header type
// 			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT , X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
// 			//Allow cross-domain settings to return to other sub-segments
// 			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last -Modified,Pragma,FooBar") //cross-domain key settings so that the browser can parse
// 			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                            //Cache request information unit is second
// 			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                   //Whether cross-domain requests need to carry cookie information is set to true by default
// 			c.Set("content-type", "application/json")                                                                                                                                                               //Set the return format to json
// 		}

// 		//Release all OPTIONS methods
// 		if method == "OPTIONS" {
// 			c.JSON(http.StatusOK, "Options Request!")
// 		}
// 		//Process the request
// 		c.Next() //Process the request
// 	}

// }
