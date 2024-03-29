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
	// initialize the embedding explicitly; 
	// here the embedded type serves as the field name
	co := container {
		base: base {
			num: 1,
		},
		str: "Heyo!",
	}
	// access the base’s fields directly on co
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)

	// invoke a method that was embedded from base directly on co
	fmt.Println("describe", co.describe())

	type describer interface {
		describe() string
	}
	// container implements the describer interface because it embeds base
	var d describer = co
	fmt.Println("describer:", d.describe())
}