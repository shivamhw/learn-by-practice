package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var Secret = "testdfvjolsdafjodsoi"


func loggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("Incoming %s %s", ctx.Request.RemoteAddr, ctx.Request.URL.Path)
		ctx.Next()
	}
}

func authMiddleware() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		token, err := ctx.Cookie("Authorization")
		if err != nil {
			fmt.Printf("err : %s", err.Error())
			ctx.AbortWithStatusJSON(401, gin.H{"err": "not auth"})
			return 
		}
		if token == "" {
			fmt.Printf("token is nil")
			ctx.Abort()
			return 
		}
		vToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signature method invalid")
			}
			return []byte(Secret), nil
		})
		if err != nil {
			fmt.Print("token parse failed, ", err.Error())
			ctx.Abort()
			return 
		}
		if !vToken.Valid {
			fmt.Printf("Token is not valid")
			ctx.Abort()
			return 
		}
		if claims, ok := vToken.Claims.(jwt.MapClaims); ok {
			userId := claims["id"]
			ctx.Set("userid", userId)
		} else {
			fmt.Printf("claim extraction failed")
			ctx.Abort()
			return 
		}
		log.Printf("here is the token: %+v", vToken)
		ctx.Next()
	}
}