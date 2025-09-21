package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var users map[string]UserCreate

type UserGetProfile struct {
	UserId string `json:"userid" uri:"userid" binding:"required"`
}

type UserCreate struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	DOB      string `json:"date" binding:"datetime=02-01-2006"` // DD-MM-YYYY
	Password string `json:"password" binding:"required,gte=4"`
}

func getRoutes() http.Handler {
	users = make(map[string]UserCreate)
	g := gin.Default()
	// g.Use(loggingMiddleware())
	g.Use(authMiddleware())
	g.GET("/echo", echo)
	registerUserRoutes(g)
	registerStreamRoutes(g)
	return g
}

func echo(c *gin.Context) {
	type hello struct {
		Msg string `json:"msg"`
	}
	userid := c.GetString("userid")
	e := hello{
		Msg: "hell0" + userid,
	}
	c.JSON(http.StatusOK, e)
}


func registerStreamRoutes(g *gin.Engine) {
	g.GET("/stream", func(ctx *gin.Context) {
		count := 0
		ctx.Stream(func(w io.Writer) bool {
			if count >10 {
				return false
			}
			ctx.SSEvent("msg", count)
			time.Sleep(time.Second * 2)
			count++
			return true
		})
	})
}

func registerUserRoutes(g *gin.Engine) {
	userG := g.Group("/users")

	userG.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": users})
	})

	userG.GET("/getUser", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name", "")
		if name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": "no name passed"})
			return
		}
		if val, ok := users[name]; !ok {
			ctx.JSON(404, gin.H{"msg": "user not found"})
			return
		} else {
			ctx.JSON(201, gin.H{"msg": val})
		}

	})

	// read from path param
	userG.GET("/:userid/profile", func(ctx *gin.Context) {
		var d UserGetProfile
		if err := ctx.ShouldBindUri(&d); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": "err in url can find profile"})
		}
		ctx.JSON(200, gin.H{"msg": "profile details : " + d.UserId})
	})

	// multi path param *
	userG.GET("/users/:userid/item/*params", func(ctx *gin.Context) {
		d := ctx.Param("params")
		ctx.JSON(200, gin.H{"msg": d})
	})
	// POST

	userG.POST("/create", func(ctx *gin.Context) {
		u := UserCreate{}
		if err := ctx.ShouldBind(&u); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "incorrect payload" + err.Error()})
			return
		}
		users[u.Name] = u
		ctx.JSON(201, gin.H{"msg": "user created"})
	})

	userG.GET("/:userid/getToken", func(ctx *gin.Context) {
		u := ctx.Param("userid")
		log.Printf("generating token for %s", u)
		eToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": u,
			"exp": time.Now().Add(time.Hour).Unix(),
		})
		sToken, err := eToken.SignedString([]byte(Secret))
		if err != nil {
			log.Println("signing token failed with ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		}
		ctx.SetCookie("Authorization", string(sToken), 3600 *24, "", "", false, false)
		ctx.JSON(200, gin.H{"token": sToken})
	})
}
