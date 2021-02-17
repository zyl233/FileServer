package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.LoadHTMLGlob("./tmp/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"mes": "6666",
		})
	})

	r.Run()
	//_ = http.ListenAndServe(":8080", http.FileServer(http.Dir("/")))
}
