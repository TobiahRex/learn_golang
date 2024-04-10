package async

import (
	// "sync"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func GoRoutineWithErrChan() {
	errChan := make(chan error)
	go func() {
		time.Sleep(time.Second * 1)
		errChan <- errors.New("error from 2nd GR")
	}()
	err := <-errChan
	fmt.Println("Error: ", err)
	fmt.Println("Closing main GR")
}

func GoRoutineWithSelect() {
	errChan := make(chan error)
	numChan := make(chan int)

	handleErrChan(errChan)
	handleNumChan(numChan)

	for {
		select {
		case err := <-errChan:
			fmt.Println("Error: ", err)
		case num := <-numChan:
			fmt.Println("Num: ", num)
		}
	}

}

func handleErrChan(errChan chan error) {
	go func() {
		for _, msg := range []string{"err1", "err2", "err3"} {
			time.Sleep(time.Second * time.Duration(rand.Float64()))
			errChan <- errors.New(msg)
		}
		close(errChan)
	}()
}
func handleNumChan(numChan chan int) {
	go func() {
		for i := range make([]int, 3) {
			n := rand.Intn(100)
			numChan <- n + i
			time.Sleep(time.Second * time.Duration(rand.Float64()))
		}
		close(numChan)
	}()
}
