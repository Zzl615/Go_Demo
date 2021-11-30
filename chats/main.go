package main

import "github.com/gin-gonic/gin"
import "net/http"
import "time"

type CreateInput struct {
	UserList map[string]interface{}
	UnionId  string
}

func CreateChat(c *gin.Context) {
	var pms CreateInput
	if err := c.ShouldBindBodyWith(&pms, binding.JSON); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
	}

}

func main() {
	router := gin.Default()
	rt_1 := router.Group("/chat_api")
	rt_1.GET("/get_some_info", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})
	rt_1.POST("/create_chat", CreateChat)

	// Listen and serve on 0.0.0.0:8080
	// mode1
	// r.Run(":8080")
	// mode2
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
