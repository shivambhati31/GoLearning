package main

import "testing"

func TestAdd(t *testing.T) {
	tcs := []struct {
		desc string
		int1 float64
		int2 float64
		int3 float64
	}{
		{"1", 1, 3, 4},
		{"2", -1, 3, 2},
		{"3", 1, -3, -2},
		{"4", 0, 0, 0},
		{"4", -6, -7, -13},
	}

	for _, tt := range tcs {
		t.Run(tt.desc, func(t *testing.T) {
			out := add(tt.int1, tt.int2)
			if out != tt.int3 {
				t.Error("Test Case didn't pass at description ", tt.desc)
			}
		})
	}
}
