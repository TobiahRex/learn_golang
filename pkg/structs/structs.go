package structs

import (
	"fmt"
)

func Structs() {
	goku := Saiyan{"Goku", 900, nil}
	gokuCopy := goku
	gokuCopy.Name = "Goku Copy"
	gokuCopy.Super()
	
	for key, value := range goku.Map() {
		fmt.Println("key: ", key, "| value: ", value)
	}
	for key, value := range gokuCopy.Map() {
		fmt.Println("key: ", key, "| value: ", value)
	}
}


type Saiyan struct {
	Name string
	Power int
	Father *Saiyan
}

func (s Saiyan) Map() map[string]interface{} {
	return map[string]interface{}{
		"Name": s.Name,
		"Power": s.Power,
		"Father": s.Father,
	}
}

func (s *Saiyan) Super() {
	s.Power += 10000
}

func NewSaiyanByRef(name string, power int, father *Saiyan) *Saiyan {
	return &Saiyan {
		Name: name,
		Power: power,
		Father: father,
	}
}

func NewSaiyanByVal(name string, power int, father *Saiyan) Saiyan {
	return Saiyan {
		Name: name,
		Power: power,
		Father: nil,
	}
}
