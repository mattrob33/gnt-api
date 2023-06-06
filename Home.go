package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gnt-api/gnt"
	"net/http"
)

func home(c *gin.Context) {
	var output string
	for i := 1; i <= 18; i++ {
		verse := gnt.GetVerse(3, 1, i)
		output += fmt.Sprintf("[%d] %s\n", i, verse.PlainText())
	}

	c.String(http.StatusOK, output)
}
