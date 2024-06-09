package error

import (
	"fmt"
)

type ClientError struct {
	Code    int
	Message string
	Raw     error
}

func (e ClientError) Error() string {
	return fmt.Sprintf("%d\t%s", e.Code, e.Message)
}
