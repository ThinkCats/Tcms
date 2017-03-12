package middleware

import (
	"net/http"

	"github.com/pjebs/restgate"
	"gopkg.in/gin-gonic/gin.v1"
)

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
