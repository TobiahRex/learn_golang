package iterators

import (
	"fmt"
)

func Maps() {
	fmt.Println("Create Map:")
	createMap()
	fmt.Println("Create Engineer:")
	createEngineer()
	fmt.Println("Create Sized Map:")
	createSizedMap()
	fmt.Println("Struct with Map:")
	structWithMap()
}

func createMap() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]
	fmt.Println("vegeta - power, exists: ", power, exists)
	fmt.Println("Whole Map: ", lookup)
	fmt.Println("goku: ", lookup["goku"])
	fmt.Println("len(map): ", len(lookup))
	delete(lookup, "goku")
	fmt.Println("delete 'goku' result: ", lookup)
}

func createSizedMap() {
	people := make(map[string]string, 3)
	people["key1"] = "value1"
	people["key2"] = "value2"
	people["key3"] = "value3"
	fmt.Println("people: ", people)
	people["key4"] = "value4"
	fmt.Println("people: ", people)
}

type Engineer struct {
	Name string
	Skills Skills
}

type Skill int
const (
	Java Skill = iota
	Golang
	Python
	JavaScript
	Cpp
)
type Skills map[Skill]bool

func (s Skill) String() string {
	return [...]string{"Java", "Golang", "Python", "JavaScript", "C++"}[s]
}


func createEngineer() {
	// Create an engineer with skills
	engineer := Engineer {
		Name: "John",
		Skills: Skills{
			Java: true,
			Golang: true,
			Python: true,
		},
	}
	fmt.Println("Engineer Name: ", engineer.Name)
	for skill, has := range engineer.Skills {
		fmt.Println("Skill: ", skill, "| Has: ", has)
	}
}

func structWithMap() {
	type JobsMap map[string]string

	type StrWithMap struct {
		Name string
		Jobs JobsMap
	}
	strWithMap := StrWithMap {
		Name: "Bob",
		Jobs: JobsMap{
			"job1": "Software Engineer",
			"job2": "Data Scientist",
			"job3": "Day Trader",
		},
	}
	for key, value := range strWithMap.Jobs {
		fmt.Println("key: ", key, "| value: ", value)
	}
}