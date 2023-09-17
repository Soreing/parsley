package writer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestUInt8Digits tests if uint8s return their digit count correctly
func TestUInt8Digits(t *testing.T) {
	tests := []struct {
		Name   string
		Value  uint8
		Digits int
	}{
		{
			Name:   "1 Digit",
			Value:  0,
			Digits: 1,
		},
		{
			Name:   "2 Digits",
			Value:  55,
			Digits: 2,
		},
		{
			Name:   "3 Digits",
			Value:  255,
			Digits: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res := ui8dc(test.Value)
			assert.Equal(t, test.Digits, res)
		})
	}
}

// TestUInt16Digits tests if uint16s return their digit count correctly
func TestUInt16Digits(t *testing.T) {
	tests := []struct {
		Name   string
		Value  uint16
		Digits int
	}{
		{
			Name:   "1 Digit",
			Value:  0,
			Digits: 1,
		},
		{
			Name:   "2 Digits",
			Value:  55,
			Digits: 2,
		},
		{
			Name:   "3 Digits",
			Value:  555,
			Digits: 3,
		},
		{
			Name:   "4 Digits",
			Value:  5555,
			Digits: 4,
		},
		{
			Name:   "5 Digits",
			Value:  55555,
			Digits: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res := ui16dc(test.Value)
			assert.Equal(t, test.Digits, res)
		})
	}
}

// TestUInt32Digits tests if uint32s return their digit count correctly
func TestUInt32Digits(t *testing.T) {
	tests := []struct {
		Name   string
		Value  uint32
		Digits int
	}{
		{
			Name:   "1 Digit",
			Value:  0,
			Digits: 1,
		},
		{
			Name:   "2 Digits",
			Value:  55,
			Digits: 2,
		},
		{
			Name:   "3 Digits",
			Value:  555,
			Digits: 3,
		},
		{
			Name:   "4 Digits",
			Value:  5555,
			Digits: 4,
		},
		{
			Name:   "5 Digits",
			Value:  55555,
			Digits: 5,
		},
		{
			Name:   "6 Digit",
			Value:  555555,
			Digits: 6,
		},
		{
			Name:   "7 Digits",
			Value:  5555555,
			Digits: 7,
		},
		{
			Name:   "8 Digits",
			Value:  55555555,
			Digits: 8,
		},
		{
			Name:   "9 Digits",
			Value:  555555555,
			Digits: 9,
		},
		{
			Name:   "10 Digits",
			Value:  1555555555,
			Digits: 10,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res := ui32dc(test.Value)
			assert.Equal(t, test.Digits, res)
		})
	}
}

// TestUInt64Digits tests if uint64s return their digit count correctly
func TestUInt64Digits(t *testing.T) {
	tests := []struct {
		Name   string
		Value  uint64
		Digits int
	}{
		{
			Name:   "1 Digit",
			Value:  0,
			Digits: 1,
		},
		{
			Name:   "2 Digits",
			Value:  55,
			Digits: 2,
		},
		{
			Name:   "3 Digits",
			Value:  555,
			Digits: 3,
		},
		{
			Name:   "4 Digits",
			Value:  5555,
			Digits: 4,
		},
		{
			Name:   "5 Digits",
			Value:  55555,
			Digits: 5,
		},
		{
			Name:   "6 Digit",
			Value:  555555,
			Digits: 6,
		},
		{
			Name:   "7 Digits",
			Value:  5555555,
			Digits: 7,
		},
		{
			Name:   "8 Digits",
			Value:  55555555,
			Digits: 8,
		},
		{
			Name:   "9 Digits",
			Value:  555555555,
			Digits: 9,
		},
		{
			Name:   "10 Digits",
			Value:  1555555555,
			Digits: 10,
		},
		{
			Name:   "11 Digit",
			Value:  55555555555,
			Digits: 11,
		},
		{
			Name:   "12 Digits",
			Value:  555555555555,
			Digits: 12,
		},
		{
			Name:   "13 Digits",
			Value:  5555555555555,
			Digits: 13,
		},
		{
			Name:   "14 Digits",
			Value:  55555555555555,
			Digits: 14,
		},
		{
			Name:   "15 Digits",
			Value:  555555555555555,
			Digits: 15,
		},
		{
			Name:   "16 Digit",
			Value:  5555555555555555,
			Digits: 16,
		},
		{
			Name:   "17 Digits",
			Value:  55555555555555555,
			Digits: 17,
		},
		{
			Name:   "18 Digits",
			Value:  555555555555555555,
			Digits: 18,
		},
		{
			Name:   "19 Digits",
			Value:  5555555555555555555,
			Digits: 19,
		},
		{
			Name:   "10 Digits",
			Value:  15555555555555555555,
			Digits: 20,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res := ui64dc(test.Value)
			assert.Equal(t, test.Digits, res)
		})
	}
}
