package main

import (
	"errors"
	"fmt"
)

func doubleEvenWithErrors(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}

func doubleEvenWithFmt(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number")
	}
	return i * 2, nil
}
