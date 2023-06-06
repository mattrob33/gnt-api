package api

import (
	"github.com/gin-gonic/gin"
	"gnt-api/gnt"
	"net/http"
)

func Verse(c *gin.Context) {
	book := requireParamInt(c, "book")
	chapter := requireParamInt(c, "chapter")
	verseNum := requireParamInt(c, "verse")

	verse := gnt.GetVerse(book, chapter, verseNum)

	c.JSON(http.StatusOK, verse)
}

func Verses(c *gin.Context) {
	book := requireParamInt(c, "book")
	chapter := requireParamInt(c, "chapter")
	startVerse := requireParamInt(c, "start_verse")
	endVerse := requireParamInt(c, "end_verse")

	var verses []gnt.Verse

	for verseNum := startVerse; verseNum <= endVerse; verseNum++ {
		verse := gnt.GetVerse(book, chapter, verseNum)
		verses = append(verses, verse)
	}

	c.JSON(http.StatusOK, verses)
}
