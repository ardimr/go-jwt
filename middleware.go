package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Group    string `json:"Group"`
}

func Authenticate(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()

	if !ok {
		c.AbortWithError(http.StatusUnauthorized, errors.New("this is not basic auth"))
		return
	}

	if secret, ok := users[username]; !ok || secret != password {
		c.AbortWithError(http.StatusUnauthorized, errors.New("wrong username or password"))
		return
	}

	newToken, err := GenerateNewToken(username)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusInternalServerError, errors.New("failed to generate new token"))
	}

	// c.SetCookie("token", newToken, 60, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": newToken,
	})

}

func Authorize(c *gin.Context) {
	// Check the value of bearer in header
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]
	fmt.Println(tokenString)
	if len(authHeader) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := ValidateToken(tokenString)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		fmt.Println(claims["Username"])
	} else {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}
	c.JSON(http.StatusOK, claims)
	fmt.Println("Authorized:", claims)
}

func GenerateNewToken(username string) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "MY_APPLICATION",
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
		Username: username,
		Email:    username + "@gmail.com",
		Group:    "Admin",
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	mySigningKey := []byte("ardimr")
	signedToken, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	mySigningKey := []byte("ardimr")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil

}
func Midelleware2(c *gin.Context) {
	fmt.Println("lolololo")
}
