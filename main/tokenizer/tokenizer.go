package tokenizer

import (
	"bufio"
	"os"
	"strings"
)

func New() []string {
	lexicalSeps := map[string]bool{
		" ": true,
		",": true,
		"+": true,
		"&": true,
		"(": true,
		")": true,
		"-": true,
		"|": true,
		"<": true,
		"[": true,
		"]": true,
		"*": true,
		"^": true,
		">": true,
		"{": true,
		"}": true,
		"/": true,
		"=": true,
		";": true,
		"%": true,
		"!": true,
		".": true,
		":": true,
	}

	selfFile, err := os.Open("./main.go")
	if err != nil {
		panic(err)
	}

	bufReader := bufio.NewReader(selfFile)
	ln, err := bufReader.ReadString('\n')

	tokenSlice := []string{}

	for err == nil {
		lineLength := len(ln)
		var token strings.Builder

		for i := 0; i < lineLength; i++ {
			char := ln[i]

			if lexicalSeps[string(char)] {
				tokenSlice = append(tokenSlice, token.String())
				tokenSlice = append(tokenSlice, string(char))
				token.Reset()
			} else {
				token.WriteByte(char)
			}

			if i == lineLength-1 {
				tokenSlice = append(tokenSlice, token.String())
			}
		}
		ln, err = bufReader.ReadString('\n')
	}

	return tokenSlice
}
