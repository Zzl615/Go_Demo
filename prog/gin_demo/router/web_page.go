package router

import "github.com/gin-gonic/gin"
import "net/http"
import "log"

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello, My World!",
	})
}

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func StartPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}
