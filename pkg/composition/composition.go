package composition;

import (
	"fmt"
)

func Composition() {
	a_goku := &Person{"Goku", "I'm a Saiyan!"}

	goku := Saiyan {
		a_goku,
		9000,
	}
	goku.Person.Introduce()
	a_goku.Introduction = "what's up?"
	goku.Person.Introduce()
	// goku.Introduce()
}

type Person struct {
	Name string
	Introduction string
}

type Saiyan struct {
	*Person
	power int
}

func (p *Person) Introduce() {
	fmt.Println(p.Introduction, p.Name)
}


// Overloads the Introduce method of the Person struct
func (s *Saiyan) Introduce() {
	fmt.Println("Hi, my name is", s.Name, "and my power is", s.power)
}