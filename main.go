package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/zipqt/goScrape/controllers"
	"github.com/zipqt/goScrape/database"
)

func main() {
	database.Start()

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("static/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/get", func(c *gin.Context) {
		randomIndex := rand.Intn(1250) + 1
		fmt.Println("index: " , randomIndex)
		page, err := controllers.Get(fmt.Sprint(randomIndex))
		if err != nil{
			log.Fatal(err)
		}
		c.String(200,`<img src="` + page.URL + `" alt="Random coloring page">`)
	})
	port := os.Getenv("PORT")
	portStr := fmt.Sprintf("0.0.0.0:%s",port)
	r.Run(portStr)
}
