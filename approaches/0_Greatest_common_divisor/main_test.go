package main

import "testing"

type gfcTestCase struct {
	name     string
	num1     int
	num2     int
	expected int
}

func TestGFC(t *testing.T) {
	tests := []gfcTestCase{
		{
			name:     "common divisors 12 and 20",
			num1:     12,
			num2:     20,
			expected: 4,
		},
		{
			name:     "same numbers",
			num1:     10,
			num2:     10,
			expected: 10,
		},
		{
			name:     "one is multiple of another",
			num1:     6,
			num2:     18,
			expected: 6,
		},
		{
			name:     "coprime numbers",
			num1:     7,
			num2:     9,
			expected: 1,
		},
		{
			name:     "one is one",
			num1:     1,
			num2:     99,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gfc(tt.num1, tt.num2)
			if result != tt.expected {
				t.Errorf(
					"gfc(%d, %d) = %d; expected %d",
					tt.num1,
					tt.num2,
					result,
					tt.expected,
				)
			}
		})
	}
}
