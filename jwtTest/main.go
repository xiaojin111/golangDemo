package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	tokenStr, err := Generate_JWT("1", "2")
	if err != nil {
		log.Println(err)
	}
	log.Println(tokenStr)
	var c CustomClaims
	c, err = Verify_jwt(tokenStr)
	log.Println(c.RoleId)
	log.Println(c.UserId)
}

const JwtPayloadKey = "JWT_HZ2020"

type CustomClaims struct {
	UserId string `json:"userId"`
	RoleId string `json:"roleId"`
	jwt.StandardClaims
}

//创建jwt令牌
func Generate_JWT(userId, roleId string) (tokenString string, err error) {
	claims := CustomClaims{
		userId,
		roleId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			Issuer:    "HZ",
		},
	}
	mySigningKey := []byte(JwtPayloadKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(mySigningKey)
	if err != nil {
		log.Println(err)
	}
	return
}

//验证jwt令牌
func Verify_jwt(tokenString string) (customClaims CustomClaims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtPayloadKey), nil
	})
	if token.Valid {
		customClaims.UserId = token.Claims.(jwt.MapClaims)["userId"].(string)
		customClaims.RoleId = token.Claims.(jwt.MapClaims)["roleId"].(string)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		log.Println(ve.Error())
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			err = errors.New("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			err = errors.New("Timing is everything")
		} else {
			err = errors.New("Couldn't handle this token:" + err.Error())
		}
	} else {
		err = errors.New("Couldn't handle this token:" + err.Error())
	}
	return
}
