package main

import "fmt"

func main() {
	question := "¿cómo estás?"

	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
}
