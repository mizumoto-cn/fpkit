package err

import (
	"fmt"
	"time"
)

func NewIndexOutOfRangeError(index, length int) error {
	return fmt.Errorf("gogenerics: index out of range: %d (length: %d)", index, length)
}

func NewTypeCastError(from any, to string) error {
	return fmt.Errorf("gogenerics: cannot cast type %#v to %s", from, to)
}

func NewInvaliadTimeIntarvalError(interval time.Duration) error {
	return fmt.Errorf("gogenerics: invalid time interval: %v", interval)
}
