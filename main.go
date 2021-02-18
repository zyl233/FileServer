package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	dbconn := os.Getenv("DATABASE_URL")
	getenv := os.Getenv("PORT")
	log.Println("port" + getenv)
	log.Println("psot" + dbconn)
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
