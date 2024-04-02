package iterators

import (
	"fmt"
)

func iterators() {
	for _, f := range getFuncs("Hello, World!", 2, 3, 3) {
		result, err := f()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}

	nums := []int{1, 2, 3, 4, 5}
	for i, n := range nums {
		fmt.Println("i: ", i, "| n: ", n)
	}

	people := map[string]int{
		"John": 30,
		"Bob":  25,
		"Mike": 40,
	}
	for key, value := range people {
		fmt.Println("key: ", key, "| value: ", value)
	}
}

func getFuncs(msg string, a, b, pow int) []func() (interface{}, error) {
	return []func() (interface{}, error){
		func() (interface{}, error) {
			return log(msg), nil
		},
		func() (interface{}, error) {
			return add(a, b), nil
		},
		func() (interface{}, error) {
			return power(b, pow)
		},
	}
}

func log(message string) string {
	return "Log: " + message
}

func add(a, b int) int {
	return a + b
}

func power(a, pow int) (int, error) {
	if pow == 0 {
		return 0, fmt.Errorf("power must be greater than 0")
	}
	if pow == 1 {
		return 1, nil
	}
	result := 1
	for i := 0; i < pow; i++ {
		result *= a
	}
	return result, nil
}


func Map(s []int, f func(int) int) []int {
	sm := make([]int, len(s))
	for i, v := range s {
		sm[i] = f(v)
	}
	return sm
}

func Filter(s []int, f func(int) bool) []int {
	sf := make([]int, 0)
	for _, v := range s {
		if f(v) {
			sf = append(sf, v)
		}
	}
	return sf
}

func Reduce(s []int, f func(int, int) int) int {
	sr := s[0]
	for _, v := range s[1:] {
		sr = f(sr, v)
	}
	return sr
}
