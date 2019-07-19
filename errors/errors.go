package errors

import "fmt"

type ErrorNotFound struct {
	What string
}

func (e *ErrorNotFound) Error() string {
	return fmt.Sprintf("%q was not found", e.What)
}
