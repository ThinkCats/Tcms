package middleware

import (
	"fmt"
	"net/http"
	"time"

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
	if username == "admin" && password == "admin" {
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
		c.SetCookie("token", t, 1000, "/", "/", true, true)
		c.JSON(http.StatusOK, gin.H{
			"x-token": t,
		})
	} else {
		c.JSON(401, gin.H{
			"message": "not access",
		})
	}
}

//CheckToken check token info
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("-----.....")
		token, _ := c.Cookie("token")
		fmt.Println("token from url:", token)
		t, err := jwt.ParseWithClaims(token, &LoginEntity{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if err != nil {
			fmt.Println(err)
		}
		if claims, ok := t.Claims.(*LoginEntity); ok && t.Valid {
			fmt.Println("Entity:", claims)
		} else {
			fmt.Println("Cant find entity")
			c.HTML(401, "index.tmpl", gin.H{
				"title": "not a valid token",
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
