package api

import (
	"github.com/gin-gonic/gin"
	"gnt-api/gnt"
	"net/http"
)

func Query(c *gin.Context) {
	query := gnt.Query{
		Lemma:        c.Query("lemma"),
		PartOfSpeech: c.Query("pos"),
		Person:       c.Query("person"),
		Tense:        c.Query("tense"),
		Voice:        c.Query("voice"),
		Mood:         c.Query("mood"),
		Case:         c.Query("case"),
		Number:       c.Query("number"),
		Gender:       c.Query("gender"),
	}

	verses := gnt.SendQuery(query)

	c.JSON(http.StatusOK, verses)
}
