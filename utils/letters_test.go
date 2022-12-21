package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLen(t *testing.T) {
	assert.Len(t, LettersWords, 1470)
}

func TestRefineWordSet(t *testing.T) {
	var results []string
	for _, w := range LettersWords {
		if len(w) > 9 || len(w) < 6 {
			continue
		}

		results = append(results, w)
	}

	for _, w := range results {
		fmt.Println("\"" + w + "\",")
	}
}
