package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(" > ")
	scanner.Scan()
	text := scanner.Text()

	fmt.Println("echoing: ", text)
}
