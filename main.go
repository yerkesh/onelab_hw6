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
	funcs = append(funcs, Correct)
	funcs = append(funcs, Correct)
	funcs = append(funcs, Correct)
	funcs = append(funcs, Correct2)
	funcs = append(funcs, InCorrect)
	fmt.Println(Execute(funcs, 10))
	//fmt.Println(ExecuteChan(funcs, 3, 10))
}

// InCorrect returns error number
func InCorrect() error {
	res := errors.New("error number: " + strconv.Itoa(errCount))
	errs = append(errs, res)
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
	for _, fun := range tasks {
		wg.Add(1)
			go func() {
				err := fun()
				defer wg.Done()
				mu.Lock()
				defer mu.Unlock()
				if err != nil {
					errCount++
				}
			}()
	}
	wg.Wait()
	if errCount > E {
		return errs[0]
	}
	return nil
}

//func ExecuteChan(tasks []func() error, E int) error {
//	var (
//		ch = make(chan int, 5)
//		//done = make(chan struct{})
//	)
//
//	cnt := 0
//	for _, fun := range tasks {
//		ch <- 1 // will block if there is N ints in channel
//			go func() {
//				err := fun()
//				fmt.Println(err)
//				if err != nil {
//					cnt++
//					ch <- cnt
//				}
//			}()
//		<-ch // removes an int from channel, allowing another to proceed
//	}
//	//close(ch)
//	fmt.Println(len(ch))
//	for num := range ch {
//		if num > E {
//			return errs[0]
//		}
//	}
//	return nil
//}