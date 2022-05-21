// Strings may contain escape sequences, such as \n.
// To avoid substituting \n for a new line, you can wrap text in backticks (`) instead of quotes (")
package main

import "fmt"

func main() {
	fmt.Println("peace be upon you\nupon you be peace")
	fmt.Println(`strings can span multiple lines with \n escape sequence`)
}
