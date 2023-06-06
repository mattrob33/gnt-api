package main

import (
	"github.com/gin-gonic/gin"
	"gnt-api/gnt"
	"net/http"
	"strconv"
)

func verse(c *gin.Context) {
	book, _ := strconv.Atoi(c.Query("book"))
	chapter, _ := strconv.Atoi(c.Query("chapter"))
	verseNum, _ := strconv.Atoi(c.Query("verse"))

	verse := gnt.GetVerse(book, chapter, verseNum)

	c.JSON(http.StatusOK, verse)
}
