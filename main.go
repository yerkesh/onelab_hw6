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

)
func main() {
	funcs := make([]func() error, 0, 10)
	// Creating functions which returns errors
	funcs = append(funcs, Correct)
	funcs = append(funcs, Correct)
	funcs = append(funcs, Correct)
	funcs = append(funcs, Correct)
	funcs = append(funcs, Correct)
	funcs = append(funcs, Correct2)
	funcs = append(funcs, InCorrect)
	funcs = append(funcs, InCorrect)
	funcs = append(funcs, InCorrect)
	fmt.Println(Execute(funcs,  1))
	//fmt.Println(ExecuteChan(funcs, 5,4))
}

// InCorrect returns error number
func InCorrect() error {
	res := errors.New("error number: " + strconv.Itoa(errCount))
	return res
}

func Correct() error {
	_, err := strconv.ParseFloat("5.5", 32)
	return err
}

func Correct2() error {
	_, err := strconv.ParseInt("5",10,16)
	return err
}

func Execute(tasks []func() error, E int) error {
	var errs []error
	for _, fun := range tasks {
		wg.Add(1)
			go func() {
				err := fun()
				defer wg.Done()
				mu.Lock()
				defer mu.Unlock()
				if err != nil {
					errCount++
					errs = append(errs, err)
				}
			}()
	}
	wg.Wait()
	if errCount > E {
		return errs[0]
	}
	return nil
}

func ExecuteChan(tasks []func() error, N int, E int) error {
	var (
		ch = make(chan int)
		done = make(chan struct{})
	    errs []error
	)

	for _, fun := range tasks {
			go func() {
				err := fun()
				if err != nil {
					res := errCount + 1
					sender(ch, res)
					errs = append(errs, err)
				}
			}()
	}
	close(ch)
	if receiver(ch, done, E) {
		return errs[0]
	}
	return nil
}

func sender(ch chan<- int, num int) {
	ch <- num
}

func receiver(ch <-chan int, done chan<- struct{}, E int) bool {
	for num := range ch {
		if num > E {
			return true
		}
		fmt.Println(num)
	}
	close(done)
	return false
}
