package composition;

import (
	"fmt"
)

func composition() {
	goku := Saiyan {
		&Person{"Goku"},
		9000,
	}
	goku.Introduce()
}

type Person struct {
	Name string
}

type Saiyan struct {
	*Person
	power int
}


func (p *Person) Introduce() {
	fmt.Println("Hi, my name is", p.Name)
}