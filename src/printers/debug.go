package printers

import (
	"fmt"

	"github.com/testdouble/diplomat/errors"
	"github.com/testdouble/diplomat/runners"
)

// Debug defines an unfiltered Printer.
type Debug struct{}

// Print prints all output, unfiltered.
func (t *Debug) Print(results chan runners.TestResult, errorChannel chan error) chan string {
	c := make(chan string)

	go func() {
		defer close(c)

		for result := range results {
			if result.Err != nil {
				c <- errors.Format(result.Err)
				errorChannel <- result.Err
				continue
			}

			c <- fmt.Sprintf("%v\n%v\n", result.Name, result.Diff)
		}
	}()

	return c
}
