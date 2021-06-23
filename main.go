package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
	errCount = 0
	errs []error
)
func main() {
	funcs := make([]func() error, 0, 10)
	// Creating functions which returns errors
	for i := 0; i < 5; i++ {
		funcs = append(funcs, Correct)
	}
	fmt.Println(Execute(funcs,2, 10))
}
// Correct returns error number
func Correct() error {
	res := errors.New("error number: " + strconv.Itoa(errCount))
	errs = append(errs, res)
	return res
}

func Execute(tasks []func() error, N int, E int) error {
	for _, fun := range tasks {
		wg.Add(N)
		for i := 0; i < N; i++ {
			go func() {
				fun()
				defer wg.Done()
				mu.Lock()
				defer mu.Unlock()
				errCount++
			}()
		}
	}
	wg.Wait()
	if errCount > E {
		return errs[0]
	}
	return nil
}