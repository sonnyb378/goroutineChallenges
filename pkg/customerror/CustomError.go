package customerror

import (
	"fmt"
)

type CustomError struct {
	Message string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("%v", c.Message)
}
