package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func main() {
	id := "shivam"
	secret := "tetjoij23jo4j23ioj"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"exp": time.Now(),
	})
	eToken, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	fmt.Print(eToken)
	//verify token
	vToken, err := jwt.Parse(eToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not valid signing method")
		}
		return []byte(secret), nil
	})
	if !vToken.Valid {
		fmt.Print("Token nont valid", err.Error(), reflect.TypeOf(err))
	}
	if claims, ok := vToken.Claims.(jwt.MapClaims); ok {
		gId := claims["id"].(string)
		fmt.Println("ID from token:", gId)
	}
}