package main

import (
	"errors"
	"log"
	"strconv"
	"testing"
)

func  TestExecute(t *testing.T)  {
	// Arrange
	testTable := []struct{
		tasks []func() error
		expected error
	} {
		{
			tasks: []func() error{
				func() error {
					return nil
				},
				func() error {
					return errors.New("error number: " + strconv.Itoa(errCount))
				},
		},
			expected: nil,
		},
		{
			tasks: []func() error{
				func() error {
					return errors.New("error number: " + strconv.Itoa(errCount))
				},
				func() error {
					return nil
				},
			},
			expected: nil,
		},
	}
	// Act
	for _, testCase := range testTable{
		result := Execute(testCase.tasks, 10)

		// Assert
		if result != testCase.expected {
			log.Fatalf("Incorrect result. Expect	%e, got %e", testCase.expected, result)
		}
	}
}

//func TestExecuteChan(t *testing.T) {
//	// Arrange
//	testTable := []struct{
//		tasks []func() error
//		expected error
//	} {
//		{
//			tasks: []func() error{
//				func() error {
//					res := errors.New("error number: " + strconv.Itoa(errCount))
//					errs = append(errs, res)
//					return res
//				},
//				func() error {
//					res := errors.New("error number: " + strconv.Itoa(errCount))
//					errs = append(errs, res)
//					return nil
//				},
//			},
//			expected: errors.New("error number: " + strconv.Itoa(errCount)),
//		},
//		{
//			tasks: []func() error{
//				func() error {
//					res := errors.New("error number: " + strconv.Itoa(errCount))
//					errs = append(errs, res)
//					return res
//				},
//				func() error {
//					res := errors.New("error number: " + strconv.Itoa(errCount))
//					errs = append(errs, res)
//					return nil
//				},
//			},
//			expected: errors.New("error number: " + strconv.Itoa(errCount)),
//		},
//	}
//	// Act
//	for _, testCase := range testTable{
//		result := ExecuteChan(testCase.tasks, 2, 3)
//		// Assert
//		if !reflect.DeepEqual(result, testCase.expected) {
//			log.Fatalf("Incorrect result. Expect	%e, got %e", testCase.expected, result)
//		}
//	}
//}