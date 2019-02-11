package bitarray

type BitArray struct {
	s []bool
	l uint64
}

func NewBitArray(s uint64) (BitArray, error) {
	if s == 0 {
		return BitArray{}, NewOutOfBoundsError(s)
	}

	return BitArray{make([]bool, s), s - 1}, nil
}

func (b BitArray) TurnOn(i uint64) error {
	if i == 0 || i >= b.l {
		return NewOutOfBoundsError(i)
	}

	b.s[i] = true
	return nil
}

func (b BitArray) TurnOff(i uint64) error {
	if i == 0 || i >= b.l {
		return NewOutOfBoundsError(i)
	}

	b.s[i] = false
	return nil
}

func (b BitArray) IndexValue(i uint64) (bool, error) {
	if i == 0 || i >= b.l {
		return false, NewOutOfBoundsError(i)
	}

	return b.s[i], nil
}
