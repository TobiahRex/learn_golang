package iterators

import (
	"fmt"
	"github.com/TobiahRex/learn_golang/pkg/structs"
)

func Slices() {
	fmt.Println("\nSlices initialized: ")
	slicesInitialized()

	fmt.Println("\nCreating Slices: ")
	createSlices()

	fmt.Println("\nReading Slices: ")
	readSlices()

	fmt.Println("\nWriting into Slices: ")
	saiyans := []*structs.Saiyan{
		{Name: "Goku", Power: 9000, Father: nil},
		{Name: "Gohan", Power: 3000, Father: nil},
		{Name: "Vegeta", Power: 8000, Father: nil},
	}
	results := writeIntoSlices(saiyans)
	for i, r := range results {
		fmt.Println("i: ", i, "| r: ", r)
	}

	fmt.Println("\nAppending into Slices: ")
	appended_powers := appendIntoSlices(saiyans)
	for i, p := range appended_powers {
		fmt.Println("i: ", i, "| power: ", p)
	}
}

func slicesInitialized() {
	scores := []int{1, 2, 3}
	for i, n := range scores {
		fmt.Println("i: ", i, "| n: ", n)
	}
}

func createSlices() {
	fmt.Println("Creating slices with make: ")
	scores := make([]int, 10)
	for i, n := range scores {
		fmt.Println("i: ", i, "| n: ", n)
	}

	fmt.Println("Creating slices by explicit initialization: ")
	var names = []string{"John", "Bob", "Mike"}	
	for i, n := range names {
		fmt.Println("i: ", i, "| n: ", n)
	}

	fmt.Println("Creating slice by appending")
	scores_2 := make([]int, 3)
	for i := 0; i < 10; i++ {
		scores_2 = append(scores_2, i)
	}
	for i, n := range scores_2 {
		fmt.Println("i: ", i, "| n: ", n)
	}

	var stuff []interface{}
	fmt.Println("stuff: ", stuff)
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			stuff = append(stuff, fmt.Sprintf("Name %d", i))
		} else {
			stuff = append(stuff, i)
		}
	}
	fmt.Println("stuff: ", stuff)
}

func readSlices() {
	scores := []int{1, 2, 3, 4, 5}
	for i, n := range scores {
		fmt.Println("i: ", i, "| n: ", n)
	}
	fmt.Println("scores[1:3]: ", scores[1:3])
	s := scores[3:] // Does not create a new slice, it creates a reference to the original slice.
	fmt.Println("scores reference ", s)
	s[0] = 100
	fmt.Println("s: ", s)
	fmt.Println("original scores: ", scores)
	
	s_copy := make([]int, len(scores))
	copy(s_copy, scores) // Creates a new slice with the same length as the original slice.
	fmt.Println("copy of scores: ", s_copy)
	s_copy[2] = 100
	fmt.Println("s_copy: ", s_copy)
	fmt.Println("original scores: ", scores)
}

func writeIntoSlices(saiyans []*structs.Saiyan) []int {
	powers := make([]int, len(saiyans))
	for i, s := range saiyans {
		powers[i] = s.Power
	}
	return powers
}

func appendIntoSlices(saiyans []*structs.Saiyan) []int {
	powers := make([]int, 0, len(saiyans))
	for _, s := range saiyans {
		powers = append(powers, s.Power)
	}
	return powers
}