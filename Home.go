package main

import (
	"github.com/gin-gonic/gin"
	"gnt-api/gnt"
	"net/http"
	"strconv"
)

func home(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/read")
}

func read(c *gin.Context) {
	book, err := strconv.Atoi(c.Query("book"))
	if err != nil {
		book = 0
	}

	chapter, err := strconv.Atoi(c.Query("chapter"))
	if err != nil {
		chapter = 1
	}

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
