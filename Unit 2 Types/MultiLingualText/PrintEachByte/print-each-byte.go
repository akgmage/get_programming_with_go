// Program to print each byte [ASCII character] if "shalom" one character per line
package main

import "fmt"

func main() {
	message := "shalom"
	for i := 0; i < len(message); i++ {
		c := message[i]
		fmt.Printf("%c\n", c)
	}
}
