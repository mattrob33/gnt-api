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

func verses(c *gin.Context) {
	book, _ := strconv.Atoi(c.Query("book"))
	chapter, _ := strconv.Atoi(c.Query("chapter"))
	startVerse, _ := strconv.Atoi(c.Query("start_verse"))
	endVerse, _ := strconv.Atoi(c.Query("end_verse"))

	var verses []gnt.Verse

	for verseNum := startVerse; verseNum <= endVerse; verseNum++ {
		verse := gnt.GetVerse(book, chapter, verseNum)
		verses = append(verses, verse)
	}

	c.JSON(http.StatusOK, verses)
}
