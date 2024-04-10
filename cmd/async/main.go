package main

import (
	"fmt"

	"github.com/TobiahRex/learn_golang/pkg/async"
)

func main() {
	// GoRoutineEasy()
	// GoRoutineMedium()
	GoRoutineHard()
}

func GoRoutineEasy() {
	fmt.Println("async.GoRoutineEasy: ")
	async.GoRoutineEasy()
	fmt.Println("async.GoRoutineWaitOne: ")
	async.GoRoutineWaitOne()
	fmt.Println("async.GoRoutineWaitMany: ")
	async.GoRoutineWaitMany()
}

func GoRoutineMedium() {
	fmt.Println("async.GoRoutineWithErrChan: ")
	async.GoRoutineWithErrChan()
	fmt.Println("async.GoRoutineWithSelect: ")
	async.GoRoutineWithSelect()
}

func GoRoutineHard() {
	fmt.Println("async.GoRoutinePullWithSelect: ")
	async.GoRoutinePullWithSelect()
	fmt.Println()

	fmt.Println("async.GoRoutineWorkerJobs: ")
	async.GoRoutineWorkerJobs()
}
