package main

import (
	"fmt"

	"github.com/mvaleriani/pl/main/tokenizer"
)

func main() {
	tokenSlice := tokenizer.New()
	// ast able to parse:
	// variable assignment, arithmetic ops
	fmt.Println(tokenSlice)
}
