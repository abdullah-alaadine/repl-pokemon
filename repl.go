package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" > ")
		scanner.Scan()
		text := scanner.Text()
		cleanedText := cleanInput(text)
		if len(cleanedText) == 0 {
			continue
		}
		
		fmt.Println("echoing: ", cleanedText)
	}

}

func cleanInput(str string) []string {
	lowerCase := strings.ToLower(str)
	words := strings.Fields(lowerCase)
	return words
}
