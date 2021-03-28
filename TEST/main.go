package main

import (
	"fmt"

	"github.com/limeihong-hue/123/TEST/interview"
)

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
	interview.Interview1()
}
