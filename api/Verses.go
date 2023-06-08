package api

import (
	"github.com/gin-gonic/gin"
	"gnt-api/gnt"
	"net/http"
)

func Verse(c *gin.Context) {
	book := RequireParamInt(c, "book")
	chapter := RequireParamInt(c, "chapter")
	verseNum := RequireParamInt(c, "verse")

	verse := gnt.GetVerse(book, chapter, verseNum)

	c.JSON(http.StatusOK, verse)
}

func Verses(c *gin.Context) {
	book := RequireParamInt(c, "book")
	chapter := RequireParamInt(c, "chapter")
	startVerse := RequireParamInt(c, "start_verse")
	endVerse := RequireParamInt(c, "end_verse")

	var verses []gnt.Verse

	for verseNum := startVerse; verseNum <= endVerse; verseNum++ {
		verse := gnt.GetVerse(book, chapter, verseNum)
		verses = append(verses, verse)
	}

	c.JSON(http.StatusOK, verses)
}
