package err

import (
	"fmt"
	"time"
)

func NewIndexOutOfRangeError(index, length int) error {
	return fmt.Errorf("fpkit: index out of range: [%d] with length: %d", index, length)
}

func NewTypeCastError(from any, to string) error {
	return fmt.Errorf("fpkit: cannot cast type %#v to %s", from, to)
}

func NewInvaliadTimeIntarvalError(interval time.Duration) error {
	return fmt.Errorf("fpkit: invalid time interval: [%v]", interval)
}
