package main

import (
	"fmt"
	"time"
)

type myError struct {
	When time.Time
	What string
}

func (e *myError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &myError{
		time.Now(),
		"it did work!",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
