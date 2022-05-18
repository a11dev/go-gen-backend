// routes.go

package main

import (

"github.com/a11dev/go-gen-backend/middleware"

)

func initializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(middleware.SetUserStatus())


	// Group user related routes together
	userRoutes := router.Group("/u")
	{

		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", middleware.EnsureNotLoggedIn(), performLogin)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", middleware.EnsureLoggedIn(), logout)

	}


}
