package middleware

import (
	"net/http"
	"time"

	"github.com/pjebs/restgate"
	"gopkg.in/appleboy/gin-jwt.v2"
	"gopkg.in/gin-gonic/gin.v1"
)

//CheckAdminAuth check admin auth
func CheckAdminAuth() *jwt.GinJWTMiddleware {
	authMiddleWare := &jwt.GinJWTMiddleware{
		Realm:      "TestZone",
		Key:        []byte("myKey"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
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
		Unauthorized: func(c *gin.Context, code int, msg string) {
			c.HTML(http.StatusNonAuthoritativeInfo, "401.tmpl", gin.H{
				"title": "Auth Error",
			})
		},
		TokenLookup: "header:Authorization",
	}
	return authMiddleWare
}

//CheckRestAuth check rest api auth
func CheckRestAuth(c *gin.Context) {
	rg := restgate.New("X-Auth-Key", "X-Auth-Secret", restgate.Static, restgate.Config{
		Key:                []string{"12345"},
		Secret:             []string{"secret"},
		HTTPSProtectionOff: true,
	})
	nextCalled := false
	nextAdapt := func(http.ResponseWriter, *http.Request) {
		nextCalled = true
		c.Next()
	}
	rg.ServeHTTP(c.Writer, c.Request, nextAdapt)
	if nextCalled == false {
		c.AbortWithStatus(401)
	}
}
