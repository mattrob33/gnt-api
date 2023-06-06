package gnt

import (
	"database/sql"
	"fmt"
	"gnt-api/bibleref"
	"log"
)

func GetVerseForRef(ref bibleref.BibleRef) Verse {
	return GetVerse(ref.Book, ref.Chapter, ref.Verse)
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

func SendQuery(query Query) []Verse {
	db, err := sql.Open("sqlite3", "./gnt.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	subQuery := query.ToSql()

	sqlQuery := fmt.Sprintf("SELECT * FROM text WHERE (book, chapter, verse) IN (%s)", subQuery)

	rows, err := db.Query(sqlQuery)
	defer rows.Close()

	var verses []Verse
	var words []Word

	var prevRef bibleref.BibleRef

	for rows.Next() {
		var ref bibleref.BibleRef
		var wordNum int
		var word Word

		err = rows.Scan(&ref.Book, &ref.Chapter, &ref.Verse, &wordNum, &word.Text, &word.Lemma, &word.PartOfSpeech, &word.Person, &word.Tense, &word.Voice, &word.Mood, &word.Case, &word.Number, &word.Gender)

		if (ref != prevRef) && (prevRef.Chapter != 0) {
			verses = append(verses, Verse{
				Book:    ref.Book,
				Chapter: ref.Chapter,
				Verse:   ref.Verse,
				Words:   words,
			})
			words = nil
		}

		words = append(words, word)
		prevRef = ref
	}

	return verses
}
