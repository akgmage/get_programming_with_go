/// Go supports embedding of structs and interfaces to express a more seamless composition of types
package main

import "fmt" 

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}
// ontainer embeds a base. An embedding looks like a field without a name
type container struct {
	base
	str string
}

func main(){
	co := container {
		base: base {
			num: 1,
		},
		str: "Heyo!",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe", co.describe())
}