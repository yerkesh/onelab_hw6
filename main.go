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
	ch chan int
)
func main() {
	funcs := make([]func() error, 0, 10)
	// Creating functions which returns errors
	for i := 0; i < 5; i++ {
		funcs = append(funcs, Correct)
	}
	//fmt.Println(Execute(funcs,2, 10))
	fmt.Println(ExecuteChan(funcs, 2, 11))
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

func ExecuteChan(tasks []func() error, N int, E int) error {
	ch = make(chan int)
	cnt := 0
	for _, fun := range tasks {
		for i := 0; i < N; i++ {
			go func() {
				fun()
				cnt++
				ch <- cnt
			}()
		}
	}

	for i := 0; i < len(tasks) * N; i++ {
		num := <-ch
		if num > E {
			return errs[0]
		}
	}
	close(ch)
	return nil
}