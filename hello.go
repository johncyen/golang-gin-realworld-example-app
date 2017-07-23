package main

import (
	"gopkg.in/gin-gonic/gin.v1"

    "golang-gin-starter-kit/common"
    "golang-gin-starter-kit/middlewares"
    "golang-gin-starter-kit/users"
)


func main() {

	db := common.DatabaseConnection()
	defer db.Close()

    db.DB().SetMaxIdleConns(10)
    db.AutoMigrate(&users.UserModel{})


    r := gin.Default()

	r.Use(middlewares.DatabaseMiddleware(db))

    v1 := r.Group("/api/v1")
    users.UsersRegister(v1.Group("/users"))

    v1.Use(middlewares.Auth())
    users.UserRegister(v1.Group("/user"))


    testAuth := r.Group("/api/v1/ping")

    testAuth.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

	r.Run() // listen and serve on 0.0.0.0:8080
}