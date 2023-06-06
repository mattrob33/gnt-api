package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	for i := 1; i <= 18; i++ {
		verse := GetVerse(3, 1, i)
		fmt.Printf("[%d] %s\n", i, verse.PlainText())
	}

	//router := setupRouter()
	//router.Run(getPort())
}

func GetVerse(book, chapter, verseNum int) Verse {
	db, err := sql.Open("sqlite3", "./gnt.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var words []Word

	rows, err := db.Query("SELECT `word`, `lemma`, `part_of_speech`, `person`, `tense`, `voice`, `mood`, `case`, `number`, `gender` FROM text WHERE book=? AND chapter=? AND verse=?", book, chapter, verseNum)
	defer rows.Close()

	for rows.Next() {
		var word Word
		err = rows.Scan(&word.Text, &word.Lemma, &word.PartOfSpeech, &word.Person, &word.Tense, &word.Voice, &word.Mood, &word.Case, &word.Number, &word.Gender)
		words = append(words, word)
	}

	return Verse{
		Book:    book,
		Chapter: chapter,
		Verse:   verseNum,
		Words:   words,
	}
}

type Verse struct {
	Book    int
	Chapter int
	Verse   int
	Words   []Word
}

type Word struct {
	Text         string
	Lemma        string
	PartOfSpeech string
	Person       string
	Tense        string
	Voice        string
	Mood         string
	Case         string
	Number       string
	Gender       string
}

func (verse *Verse) PlainText() string {
	var text string
	for _, word := range verse.Words {
		text += fmt.Sprintf("%s ", word.Text)
	}
	return text
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	//router.GET("/search", search)

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
