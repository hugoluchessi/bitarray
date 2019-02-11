package bitarray

import "fmt"

type OutOfBoundsError struct {
	i uint64
}

func NewOutOfBoundsError(i uint64) OutOfBoundsError {
	return OutOfBoundsError{i}
}

func (err OutOfBoundsError) Error() string {
	return fmt.Sprintf("Index [%d] out of bounds", err.i)
}
