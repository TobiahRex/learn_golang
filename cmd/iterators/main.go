package main

import (
	"fmt"
	"github.com/TobiahRex/learn_golang/pkg/iterators"
)

func main() {
	fmt.Println("Arrays: ")
	iterators.Arrays()
	fmt.Println("\nSlices: ")
	iterators.Slices()
	fmt.Println("\nStrings: ")
	iterators.Strings()
	fmt.Println("\nMaps: ")
	iterators.Maps()
}