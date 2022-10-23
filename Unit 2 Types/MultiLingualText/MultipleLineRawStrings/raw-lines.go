// Unlike conventional string literals, raw string literals can span multiple lines of source code
// The following will produce output, including the tabs used for indentation
package main

import "fmt"

func main() {
	fmt.Println(`
		peace be upon you
		upon you be peace
	`)
}
