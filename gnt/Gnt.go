package gnt

import (
	"database/sql"
	"log"
)

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
