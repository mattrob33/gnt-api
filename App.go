package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"gnt-api/api"
	"os"
)

func main() {
	router := setupRouter()
	router.Run(getPort())
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLFiles("./templates/verses.tmpl", "./templates/wordsearch.tmpl")

	router.GET("/", home)
	router.GET("/read", read)
	router.GET("/search", wordsearch)

	router.GET("/api/verse", api.Verse)
	router.GET("/api/search", api.Verses)
	router.GET("/api/query", api.Query)

	return router
}

// Get the Port from the environment so we can run on Heroku
func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
