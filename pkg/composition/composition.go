package composition;

import (
	"fmt"
)

func Composition() {
	goku := Saiyan {
		&Person{"Goku"},
		9000,
	}
	goku.Person.Introduce()
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


// Overloads the Introduce method of the Person struct
func (s *Saiyan) Introduce() {
	fmt.Println("Hi, my name is", s.Name, "and my power is", s.power)
}