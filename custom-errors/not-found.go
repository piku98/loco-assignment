package customerrors

import "fmt"

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}
