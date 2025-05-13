package main

import (
	"fmt"
	"os"
	"lexer/internal/lexer"
)

func main(){
	fileName := "./examples/00.lang"
	bytes,ok := os.ReadFile(fileName)
	if (ok !=nil) {
		fmt.Printf("File not exists %s",fileName)
		return
	}
	source := string(bytes)
	tokens := lexer.Tokenize(source)
	for _,token := range tokens {
		token.Debug()
	}
}