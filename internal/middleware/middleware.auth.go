package middleware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/shaj13/go-guardian/auth"

	"github.com/a11dev/go-gen-backend/internal/config"
)

var identityKey = "id"
var userName = "user"

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func JwtAuth(authenticator auth.Authenticator, cfg *config.Config) (*jwt.GinJWTMiddleware, error) {
	timeout, err := time.ParseDuration(cfg.JwtConfig.JWTExpiration)
	if err != nil {
		panic(err)
	}
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       cfg.JwtConfig.JWTRealm,
		Key:         []byte(cfg.JwtConfig.JWTSigningKey),
		Timeout:     timeout,
		MaxRefresh:  time.Hour,
		IdentityKey: cfg.JwtConfig.JWTId,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(auth.Info); ok {
				return jwt.MapClaims{
					identityKey: v.ID(),
					userName:    v.UserName(),
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},

		Authenticator: func(c *gin.Context) (interface{}, error) {
			if !cfg.LdapConfig.LdapEnabled {

				return auth.NewDefaultUser("SO00050", "SO00050", nil, nil), nil

			}
			user, err := authenticator.Authenticate(c.Request)
			if err == nil {
				return user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// logiche autorizzative
			return true

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
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

}
