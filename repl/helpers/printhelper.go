package helpers

import (
	"fmt"
	"time"
)

const (
	interval = 100 * time.Millisecond
)

// Should use this but wont cos it's a lot of job, using this with like a
// print ch where this function just prints with the right interval
// whenever the channel gets something would be sick tho.
func PrintArray(arr []string) {
	for _, val := range arr {
		time.Sleep(interval)
		fmt.Println(val)
	}
}
