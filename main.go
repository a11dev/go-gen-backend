package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/a11dev/go-gen-backend/internal/config"
	"github.com/a11dev/go-gen-backend/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/auth/strategies/ldap"
	"github.com/shaj13/go-guardian/store"
)

var router *gin.Engine
var authenticator auth.Authenticator
var cache store.Cache

func main() {

	// load application configurations
	cfg, err := config.Load()
	if err != nil {
		log.Printf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	setupGoGuardian(cfg)

	// Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)

	// Creates a router without any middleware by default
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Initialize the routes
	routes.InitializeRoutes(router, cfg, authenticator, cache)

	// Start serving the application
	router.RunTLS(":"+cfg.ServerPort, "./keys/server-cert.pem", "./keys/server-key.pem")

	// router.Run()
}

// Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication.
// Here we set up an ldap authentication
func setupGoGuardian(cfg *config.Config) {
	conf := &ldap.Config{
		BaseDN: cfg.LdapConfig.LdapBase,
		Port:   "389",
		Host:   cfg.LdapConfig.LdapServer,
		Filter: "(uid=%s)",
	}
	authenticator = auth.New()
	cache = store.NewFIFO(context.Background(), time.Minute*10)
	strategy := ldap.NewCached(conf, cache)
	authenticator.EnableStrategy(ldap.StrategyKey, strategy)
}
