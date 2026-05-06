package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Occupation int

const (
	Student Occupation = iota
	Teacher
	Staff
)

type user struct {
	ID         string     `json:id`
	Name       string     `json:name`
	Surname    string     `json:surname`
	Patronym   string     `json:patronym`
	Occupation Occupation `json:occupation`
}

var users = []user{}

func getUsers(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, users)
}

func getUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, a := range users {
		if a.ID == id {
			// ctx.IndentedJSON(http.StatusOK, a)
			ctx.HTML(http.StatusOK, "account.tmpl", gin.H{
				"ID":         a.ID,
				"Name":       a.Name,
				"Surname":    a.Surname,
				"Patronym":   a.Patronym,
				"Occupation": a.Occupation,
			})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func postUsers(ctx *gin.Context) {
	var newUser user
	if err := ctx.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)
	ctx.IndentedJSON(http.StatusCreated, newUser)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("static/templates/*")
	router.StaticFS("stylesheets", http.Dir("static/stylesheets"))
	router.StaticFS("images", http.Dir("static/images"))

	{
		api := router.Group("/api")
		api.GET("/ping", func(ctx *gin.Context) {
			// Return JSON response
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/users", getUsers)
		api.POST("/users", postUsers)
	}

	router.GET("/users/:id", getUserByID)

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	router.GET("/my-account", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "account.tmpl", gin.H{})
	})

	router.GET("journal", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "journal.tmpl", gin.H{})
	})

	router.Run(":3257")
}
