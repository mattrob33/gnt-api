package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gnt-api/api"
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

func read(c *gin.Context) {
	book := api.RequireParamInt(c, "book")
	chapter := api.RequireParamInt(c, "chapter")
	verses := gnt.GetChapter(book, chapter)

	bookTitles := []string{
		"Matt",
		"Mark",
		"Luke",
		"John",
		"Acts",
		"Rom",
		"1 Cor",
		"2 Cor",
		"Gal",
		"Eph",
		"Phil",
		"Col",
		"1 Thess",
		"2 Thess",
		"1 Tim",
		"2 Tim",
		"Titus",
		"Phlm",
		"Heb",
		"James",
		"1 Pet",
		"2 Pet",
		"1 John",
		"2 John",
		"3 John",
		"Jude",
		"Rev",
	}

	c.HTML(http.StatusOK, "verses.tmpl", gin.H{
		"Verses":     verses,
		"BookTitles": bookTitles,
	})
}

func wordsearch(c *gin.Context) {
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

	bookTitles := []string{
		"Matt",
		"Mark",
		"Luke",
		"John",
		"Acts",
		"Rom",
		"1 Cor",
		"2 Cor",
		"Gal",
		"Eph",
		"Phil",
		"Col",
		"1 Thess",
		"2 Thess",
		"1 Tim",
		"2 Tim",
		"Titus",
		"Phlm",
		"Heb",
		"James",
		"1 Pet",
		"2 Pet",
		"1 John",
		"2 John",
		"3 John",
		"Jude",
		"Rev",
	}

	c.HTML(http.StatusOK, "wordsearch.tmpl", gin.H{
		"Verses":     verses,
		"BookTitles": bookTitles,
	})
}
