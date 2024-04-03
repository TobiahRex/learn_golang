package iterators

import (
	"fmt"
)

func Arrays() {
	arraysWithSize()
	arraysInitialized()
}

func arraysWithSize() {
	var scores [10]int
	scores[0] = 10

	fmt.Println(scores)
}

func arraysInitialized() {
	var scores = [4]int{10, 20, 30, 40}
	fmt.Println(scores)
}