# Calcular using Pratt's Top Down Operator Precedence Algorithm

A calculator that parses expressions using pratt's top down operator precedence algorithm

This is an implementation of Pratt algorithm for parsing mathematical expressions, providing a flexible and efficient way to handle different operators and precedence levels.

## Parse Tree Diagrams

This implementation uses Pratt's Algorithm to ensure the correct parse trees are constructed for expressions, taking into consideration the different precedence levels of operators:

### Addition and Multiplication

Consider the expression "1 + 2 * 3". The parse tree for this expression can be represented as follows:

```
   +
  / \
 1   *
    / \
   2   3
```

This illustrates how the Pratt algorithm builds a parse tree, giving higher precedence to the multiplication operation.

### Exponentiation and Addition

Now, let's explore the expression "2^3 + 4". The parse tree for this expression looks like:

```
   +
  / \
 ^   4
/ \
2   3
```

This shows how the Pratt algorithm handles the exponentiation operator, which has higher precedence than addition.

<!-- ## Features

- Support for basic arithmetic operations: addition, subtraction, multiplication, division.
- Exponentiation operator (`^`) for calculating powers.
- Easily extensible for adding new operators.
- Error handling for undefined variables and other parsing errors. -->

<!-- ## Installation

To use this calculator in your Golang project, you can import it as a package:

```go
import "github.com/yourusername/calculator"
```

If you haven't already, initialize your Go module:

```bash
go mod init yourmodule
```

Then run:

```bash
go get github.com/yourusername/calculator
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/yourusername/calculator"
)

func main() {
	// Create a parser
	parser := calculator.CreateParser("3 + 4 * 2 / (1 - 5)^2")

	// Parse and calculate the result
	result := parser.Calculate()

	fmt.Println("Result:", result)
}
``` -->

## Test

```
go test -v
```

<!-- ## Contributing

If you'd like to contribute to the development of this calculator, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. -->