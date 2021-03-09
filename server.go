package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

//get method
func GetMethod(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Get successfull",
	})
}

//get method passing parameters
func PathParameters(ctx *gin.Context) {
	name := ctx.Param("name")
	age := ctx.Param("age")

	ctx.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

//post method
func PostMethod(ctx *gin.Context) {
	body := ctx.Request.Body
	value, _ := ioutil.ReadAll(body)

	ctx.JSON(200, gin.H{
		"message": string(value),
	})
}
func main() {
	Router = gin.Default()
	api := Router.Group("/api")
	{
		api.GET("/GET", GetMethod)
		api.POST("/", PostMethod)
		api.GET("/path/:name/:age", PathParameters) // /path/name/32
	}

	Router.Run(":5000")
}
