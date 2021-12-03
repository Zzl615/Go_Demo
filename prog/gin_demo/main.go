package main

import "github.com/gin-gonic/gin"
import "net/http"
import "time"
import "github.com/zzl615/gin_demo/router"

func main() {
	router_1 := gin.Default()
	router_1.LoadHTMLGlob("templates/*")
	router_1.GET("/index", router.Index)
	router_1.GET("/start_page", router.StartPage)
	
	server := &http.Server{
		Addr:           "localhost:8080",
		Handler:        router_1,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()

}
