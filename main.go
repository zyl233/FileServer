package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	dbconn := os.Getenv("DATABASE_URL")
	log.Println(dbconn)
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.LoadHTMLGlob("./tmp/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"mes": "23333",
		})
	})

	r.Run()
	//_ = http.ListenAndServe(":8080", http.FileServer(http.Dir("/")))
}
