package gnt

import "fmt"

type Query struct {
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

func (q *Query) ToSql() string {
	var conditions []string

	addCondition(&conditions, q.Lemma, "lemma")
	addCondition(&conditions, q.PartOfSpeech, "part_of_speech")
	addCondition(&conditions, q.Person, "person")
	addCondition(&conditions, q.Tense, "tense")
	addCondition(&conditions, q.Voice, "voice")
	addCondition(&conditions, q.Mood, "mood")
	addCondition(&conditions, q.Case, "case")
	addCondition(&conditions, q.Number, "number")
	addCondition(&conditions, q.Gender, "gender")

	var sql = "SELECT `book`, `chapter`, `verse` FROM text WHERE "

	for i := 0; i < len(conditions); i++ {
		if i != 0 {
			sql += " AND"
		}
		sql += conditions[i]
	}

	return sql
}

func addCondition(conditions *[]string, param, column string) {
	if param != "" {
		*conditions = append(*conditions, fmt.Sprintf(" (`%s`='%s')", column, param))
	}

}
