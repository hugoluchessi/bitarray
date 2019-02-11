package bitarray_test

import (
	"fmt"
	"testing"

	"github.com/hugoluchessi/bitarray"
)

func TestOutOfBoundsErrorString(t *testing.T) {
	var idx uint64
	idx = 10
	err := bitarray.NewOutOfBoundsError(idx)
	expectedMsg := fmt.Sprintf("Index [%d] out of bounds", idx)
	msg := err.Error()

	if msg != expectedMsg {
		t.Errorf("Expected message to be '%s' got '%s'", expectedMsg, msg)
	}

	t.Log("SUCCESS")
}
func TestNewBitArray(t *testing.T) {
	t.Run("Correct Parameter", func(t *testing.T) {
		var s uint64 = 200
		_, err := bitarray.NewBitArray(s)

		if err != nil {
			t.Error("Must not return error")
		}
	})

	t.Run("Incorrect Parameter", func(t *testing.T) {
		var s uint64 = 0
		_, err := bitarray.NewBitArray(s)

		if err == nil {
			t.Error("Must return error")
		}
	})
}

func TestTurnOn(t *testing.T) {
	var s uint64 = 10
	bs, err := bitarray.NewBitArray(s)

	t.Run("Correct Parameter", func(t *testing.T) {
		err = bs.TurnOn(8)

		if err != nil {
			t.Error("Error should be nil.")
		}
	})

	t.Run("Index equals the set length", func(t *testing.T) {
		err = bs.TurnOn(s)

		if err == nil {
			t.Error("Error should be returned")
		}
	})

	t.Run("Index greater than set length", func(t *testing.T) {
		err = bs.TurnOn(s + 1)

		if err == nil {
			t.Error("Error should be returned")
		}
	})
}

func TestTurnOff(t *testing.T) {
	var s uint64 = 10
	bs, err := bitarray.NewBitArray(s)

	t.Run("Correct Parameter", func(t *testing.T) {
		err = bs.TurnOff(8)

		if err != nil {
			t.Error("Error should be nil.")
		}
	})

	t.Run("Index equals the set length", func(t *testing.T) {
		err = bs.TurnOff(s)

		if err == nil {
			t.Error("Error should be returned")
		}
	})

	t.Run("Index greater than set length", func(t *testing.T) {
		err = bs.TurnOff(s + 1)

		if err == nil {
			t.Error("Error should be returned")
		}
	})
}

func TestIndexValue(t *testing.T) {
	var s uint64 = 10

	t.Run("Correct Parameter", func(t *testing.T) {
		bs, err := bitarray.NewBitArray(s)
		var idx uint64 = 8
		v, err := bs.IndexValue(idx)

		if err != nil {
			t.Error("Error should be nil.")
		}

		if v {
			t.Error("Value should be false.")
		}

		err = bs.TurnOn(idx)
		v, err = bs.IndexValue(idx)

		if err != nil {
			t.Error("Error should be nil.")
		}

		if !v {
			t.Error("Value should be true.")
		}

		err = bs.TurnOff(idx)
		v, err = bs.IndexValue(idx)

		if err != nil {
			t.Error("Error should be nil.")
		}

		if v {
			t.Error("Value should be false.")
		}
	})

	t.Run("Index equals the set length", func(t *testing.T) {
		bs, err := bitarray.NewBitArray(s)
		_, err = bs.IndexValue(s)

		if err == nil {
			t.Error("Error should be returned")
		}
	})

	t.Run("Index greater than set length", func(t *testing.T) {
		bs, err := bitarray.NewBitArray(s)
		_, err = bs.IndexValue(s + 1)

		if err == nil {
			t.Error("Error should be returned")
		}
	})
}
