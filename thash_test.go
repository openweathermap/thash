package thash

import (
	"testing"
)

func TestMaxZoom(t *testing.T) {
	tests := []struct {
		h   int64
		out int
	}{
		{1, 1},
		{12, 2},
		{123, 3},
		{1234, 4},
	}

	for _, test := range tests {
		got := MaxZoom(test.h)

		if got != test.out {
			t.Fatalf("MaxZoom(%d): %d; want: %d", test.h, got, test.out)
		}
	}
}

func TestGetDigit(t *testing.T) {
	tests := []struct {
		h   int64
		r   int
		out int
	}{
		{1, 1, 1},
		{12, 1, 1},
		{12, 2, 2},
		{123, 1, 1},
		{123, 2, 2},
		{123, 3, 3},
	}

	for _, test := range tests {
		got := getDigit(test.h, test.r)

		if got != test.out {
			t.Fatalf("getDigit(%d, %d): %d; want: %d", test.h, test.r, got, test.out)
		}
	}
}

func TestZXYtoHash(t *testing.T) {
	tests := []struct {
		in  [3]int
		out int64
	}{
		{[3]int{1, 0, 0}, 1},
		{[3]int{1, 1, 0}, 2},
		{[3]int{1, 0, 1}, 3},
		{[3]int{1, 1, 1}, 4},

		{[3]int{2, 0, 0}, 11},
		{[3]int{2, 1, 0}, 12},
		{[3]int{2, 2, 0}, 21},
		{[3]int{2, 3, 0}, 22},
		{[3]int{2, 0, 3}, 33},
		{[3]int{2, 1, 3}, 34},
		{[3]int{2, 2, 3}, 43},
		{[3]int{2, 3, 3}, 44},
	}

	for _, test := range tests {
		in := test.in
		got := ZXYtoHash(in[0], in[1], in[2])

		if got != test.out {
			t.Fatalf("ZXYtoHash(%d, %d, %d): %d; want: %d", in[0], in[1], in[2], got, test.out)
		}
	}
}

func TestHashtoZXY(t *testing.T) {
	tests := []struct {
		in  int64
		out [3]int
	}{
		{1, [3]int{1, 0, 0}},
		{2, [3]int{1, 1, 0}},
		{3, [3]int{1, 0, 1}},
		{4, [3]int{1, 1, 1}},

		{11, [3]int{2, 0, 0}},
		{12, [3]int{2, 1, 0}},
		{21, [3]int{2, 2, 0}},
		{22, [3]int{2, 3, 0}},
		{33, [3]int{2, 0, 3}},
		{34, [3]int{2, 1, 3}},
		{43, [3]int{2, 2, 3}},
		{44, [3]int{2, 3, 3}},
	}

	for _, test := range tests {
		out := test.out
		z, x, y := HashtoZXY(test.in)

		if z != out[0] || x != out[1] || y != out[2] {
			t.Fatalf("HashtoZXY(%d): %d, %d, %d; want: %d, %d, %d", test.in, z, x, y, out[0], out[1], out[2])
		}
	}
}

func TestCentralPoint(t *testing.T) {
	tests := []struct {
		in  int64
		out [2]float32
	}{
		{1, [2]float32{66.51326, -90}},
		{2, [2]float32{66.51326, 90}},
	}

	for _, test := range tests {
		got := CentralPoint(test.in)

		if got[0] != test.out[0] || got[1] != test.out[1] {
			t.Fatalf("CentralPoint(%d): %v; want: %v", test.in, got, test.out)
		}
	}
}
