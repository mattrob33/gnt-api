package gnt

import "fmt"

type Verse struct {
	Book    int
	Chapter int
	Verse   int
	Words   []Word
}

func (verse *Verse) PlainText() string {
	var text string
	for _, word := range verse.Words {
		text += fmt.Sprintf("%s ", word.Text)
	}
	return text
}
