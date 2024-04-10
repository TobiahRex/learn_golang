package async

import (
	"fmt"
	"sync"
	"time"
)

func GoRoutineEasy() {
	go func() {
		fmt.Println("2nd GR")
		takeTime()
		fmt.Println("2nd GR done")
	}()
	time.Sleep(time.Second * 1)
	fmt.Println("Closing main GR")
}
func GoRoutineWaitOne() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("2nd GR")
		takeTime()
		fmt.Println("2nd GR done")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Closing main GR")
}

func GoRoutineWaitMany() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	for _, msg := range []string{"1st GR", "2nd GR", "3rd GR"} {
		go func(m string) {
			fmt.Println("Start ", m)
			takeTime()
			fmt.Println("End ", m)
			wg.Done()
		}(msg)
	}
	wg.Wait()
	fmt.Println("Closing main GR")
}

func takeTime() {
	time.Sleep(time.Second * 1)
}
