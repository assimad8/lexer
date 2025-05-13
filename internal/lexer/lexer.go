package lexer

import (
	"os" 
	"fmt"
	"regexp"
)





func Source(fileName string) {
	bytes,ok := os.ReadFile(fileName)
	if (ok !=nil) {
		fmt.Printf("File not exists %s",fileName)
		return
	}
	source := string(bytes)
	fmt.Printf("Code: %s\n",source)
}