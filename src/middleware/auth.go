package middleware

import (
	"fmt"
	"net/http"
	"time"

	"tcms/src/dao"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pjebs/restgate"
	"gopkg.in/gin-gonic/gin.v1"
)

//LoginEntity Login form entity
type LoginEntity struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	jwt.StandardClaims
}

//CheckLogin check user login and set token info
func CheckLogin(c *gin.Context) {
	//TODO check username and password
	username := c.PostForm("username")
	password := c.PostForm("password")
	user, err := dao.QueryUser(username)
	if err != nil {
		c.HTML(401, "login.tmpl", gin.H{
			"message": "login error, username may not existed",
		})
	}
	if password == user.Password {
		claims := &LoginEntity{
			"admin",
			"123456",
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			c.JSON(401, gin.H{
				"message": "error",
			})
		}
		cookie := &http.Cookie{
			Name:   "token",
			Value:  t,
			Path:   "/",
			MaxAge: 1000,
		}
		http.SetCookie(c.Writer, cookie)
		http.Redirect(c.Writer, c.Request, "/admin", 302)
	} else {
		c.HTML(401, "login.tmpl", gin.H{
			"message": "login error",
		})
	}
}

//CheckToken check token info
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")
		t, err := jwt.ParseWithClaims(token, &LoginEntity{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if err != nil {
			c.HTML(401, "login.tmpl", gin.H{
				"message": "not a valid token",
			})
		}
		if claims, ok := t.Claims.(*LoginEntity); ok && t.Valid {
			fmt.Println("Entity:", claims)
		} else {
			c.HTML(401, "login.tmpl", gin.H{
				"message": "not a valid token",
			})
			c.Abort()
		}
		c.Next()
	}
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
