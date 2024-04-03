package iterators

import (
	"fmt"
)

func Slices() {
	fmt.Println("\nSlices initialized: ")
	slicesInitialized()

	fmt.Println("\nCreating Slices: ")
	createSlices()

	fmt.Println("\nReading Slices: ")
	readSlices()
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
	s := scores[:]
	fmt.Println("scores copy ", s)
	s[0] = 100
	fmt.Println("s: ", s)
}
