package sample

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("JWT")

type jwtAuthClaims struct {
	jwt.StandardClaims
	UserId string `json:"userId"`
}

func demoJwt(ginEngine *gin.Engine) {
	ginEngine.GET("/test", func(ginContext *gin.Context) {
		tokenString, err := jwtGenerate("admin")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(tokenString)
			{
				claims, err := jwtVerify(tokenString)
				fmt.Printf("%#v\n", claims)
				fmt.Println(err)
			}
			{
				claims, err := jwtVerify(tokenString + "1")
				fmt.Printf("%#v\n", claims)
				fmt.Println(err)
			}
			{
				claims, err := jwtVerify("Header.Payload.Signature")
				fmt.Printf("%#v\n", claims)
				fmt.Println(err)
			}
			time.Sleep(2 * time.Second)
			{
				claims, err := jwtVerify(tokenString)
				fmt.Printf("%#v\n", claims)
				fmt.Println(err)
			}
		}
		ginContext.Status(http.StatusOK)
	})
	ginEngine.GET("/create", func(ginContext *gin.Context) {
		token, err := jwtGenerate("admin")
		if err == nil {
			ginContext.String(http.StatusOK, token)
		} else {
			ginContext.String(http.StatusBadRequest, err.Error())
		}
	})
	ginEngine.GET("/verify", func(ginContext *gin.Context) {
		value := ginContext.GetHeader("Authorization")
		claims, err := jwtVerify(value)
		fmt.Printf("%#v\n", claims)
		fmt.Println(err)
		if err == nil {
			ginContext.Status(http.StatusOK)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
}

func jwtGenerate(userId string) (string, error) {
	// expiresAt := time.Now().Add(24 * time.Hour).Unix()
	expiresAt := time.Now().Add(1 * time.Second).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwtAuthClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userId,
			ExpiresAt: expiresAt,
		},
		UserId: userId,
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func jwtVerify(tokenString string) (*jwtAuthClaims, error) {
	var claims jwtAuthClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if token == nil {
		return nil, err
	} else {
		if !token.Valid {
			if strings.HasPrefix(err.Error(), "token is expired by") {
				return &claims, err
			} else {
				return nil, err
			}
		} else {
			return &claims, nil
		}
	}
}
