package main

import (
	"bufio"
	"fmt"
	"os"
)

const PROMPT = ">> "

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input a mathematical expression to evaluate e.g 2 ^ 3 * 4 + 1")

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		lexer := NewLexer(line)
		parser := NewParser(lexer)
		calculator := parser.Parse()
		parser.LogErrors()

		result := calculator.Evaluate()

		fmt.Printf("Result: %f\n", result)
	}
}
