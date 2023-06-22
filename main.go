package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Articles struct {
	Name string
	Tags []string
}

func main() {
	articles := getArticlesNames()

	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("template/*/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "template/pages/index.html", nil)
	})
	router.GET("/articles", func(c *gin.Context) {
		c.HTML(200, "template/pages/articles.html", gin.H{
			"articles": articles,
		})
	})
	router.GET("/carrer", func(c *gin.Context) {
		c.HTML(200, "template/pages/carrer.html", nil)
	})
	router.GET("/skills", func(c *gin.Context) {
		c.HTML(200, "template/pages/skills.html", nil)
	})
  router.GET("/about", func(c *gin.Context) {
		c.HTML(200, "template/pages/about.html", nil)
	})
	err := router.Run(":9999")
	if err != nil {
		log.Fatal(err)
	}
}

func getArticlesNames() []string {
	files, err := os.ReadDir("template/articles/")
	if err != nil {
		log.Fatal(err)
	}

	var sliceFiles []string

	for _, file := range files {
		sliceFiles = append(sliceFiles, file.Name())
	}

	return sliceFiles
}
