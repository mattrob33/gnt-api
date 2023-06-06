package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"gnt-api/gnt"
	"net/http"
	"os"
)

func main() {
	router := setupRouter()
	router.Run(getPort())
}

func home(c *gin.Context) {
	var output string
	for i := 1; i <= 18; i++ {
		verse := gnt.GetVerse(3, 1, i)
		output += fmt.Sprintf("[%d] %s\n", i, verse.PlainText())
	}

	c.String(http.StatusOK, output)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", home)

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
