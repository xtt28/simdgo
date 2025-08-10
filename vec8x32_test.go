package simdgo_test

import (
	"reflect"
	"testing"

	"github.com/xtt28/simdgo"
)

func TestVec8x32Add(t *testing.T) {
	tests := []struct {
		a        [8]float32
		b        [8]float32
		expected [8]float32
	}{
		{
			[...]float32{1, 2, 3, 4, 5, 6, 7, 8},
			[...]float32{5, 6, 7, 8, 9, 10, 11, 12},
			[...]float32{6, 8, 10, 12, 14, 16, 18, 20},
		},
	}

	for _, test := range tests {
		dst := &[8]float32{}
		simdgo.Vec8x32Add(&test.a, &test.b, dst)

		if !reflect.DeepEqual(*dst, test.expected) {
			t.Fatalf("Vec8x32Add(%v, %v): expected %v, got %v", test.a, test.b, test.expected, *dst)
		}
	}
}

func TestVec8x32Sub(t *testing.T) {
	tests := []struct {
		a        [8]float32
		b        [8]float32
		expected [8]float32
	}{
		{
			[...]float32{5.4, 6, 7, 8, 9, 7, 1, 4},
			[...]float32{1, 5, 2, 9, 3, 2, 1, 0},
			[...]float32{4.4, 1, 5, -1, 6, 5, 0, 4},
		},
	}

	for _, test := range tests {
		dst := &[8]float32{}
		simdgo.Vec8x32Sub(&test.a, &test.b, dst)

		if !reflect.DeepEqual(*dst, test.expected) {
			t.Fatalf("Vec8x32Sub(%v, %v): expected %v, got %v", test.a, test.b, test.expected, *dst)
		}
	}
}

func TestVec8x32Mul(t *testing.T) {
	tests := []struct {
		a        [8]float32
		b        [8]float32
		expected [8]float32
	}{
		{
			[...]float32{5.4, 6, 7, 8, 1, 2, 3, 4},
			[...]float32{1, 5, 2, 9, 2, 4, 6, 8},
			[...]float32{5.4, 30, 14, 72, 2, 8, 18, 32},
		},
	}

	for _, test := range tests {
		dst := &[8]float32{}
		simdgo.Vec8x32Mul(&test.a, &test.b, dst)

		if !reflect.DeepEqual(*dst, test.expected) {
			t.Fatalf("Vec8x32Mul(%v, %v): expected %v, got %v", test.a, test.b, test.expected, *dst)
		}
	}
}

func TestVec8x32Div(t *testing.T) {
	tests := []struct {
		a        [8]float32
		b        [8]float32
		expected [8]float32
	}{
		{
			[...]float32{5.4, 6, 7, 8, 2, 6, 24, 35},
			[...]float32{2, 3, 2, 5, 2, 2, 6, 7},
			[...]float32{2.7, 2, 3.5, 1.6, 1, 3, 4, 5},
		},
	}

	for _, test := range tests {
		dst := &[8]float32{}
		simdgo.Vec8x32Div(&test.a, &test.b, dst)

		if !reflect.DeepEqual(*dst, test.expected) {
			t.Fatalf("Vec8x32Div(%v, %v): expected %v, got %v", test.a, test.b, test.expected, *dst)
		}
	}
}

func BenchmarkVec8x32Add(b *testing.B) {
	x := &[8]float32{1, 2, 3, 4, 5, 6, 7, 8}
	y := &[8]float32{5, 6, 7, 8, 9, 5, 6, 8}
	dst := &[8]float32{}
	for b.Loop() {
		simdgo.Vec8x32Add(x, y, dst)
	}
}

func BenchmarkVec8x32Sub(b *testing.B) {
	x := &[8]float32{1, 2, 3, 4, 5, 6, 7, 8}
	y := &[8]float32{5, 6, 7, 8, 9, 5, 6, 8}
	dst := &[8]float32{}
	for b.Loop() {
		simdgo.Vec8x32Sub(x, y, dst)
	}
}

func BenchmarkVec8x32Mul(b *testing.B) {
	x := &[8]float32{1, 2, 3, 4, 5, 6, 7, 8}
	y := &[8]float32{5, 6, 7, 8, 9, 5, 6, 8}
	dst := &[8]float32{}
	for b.Loop() {
		simdgo.Vec8x32Mul(x, y, dst)
	}
}

func BenchmarkVec8x32Div(b *testing.B) {
	x := &[8]float32{1, 2, 3, 4, 5, 6, 7, 8}
	y := &[8]float32{5, 6, 7, 8, 9, 5, 6, 8}
	dst := &[8]float32{}
	for b.Loop() {
		simdgo.Vec8x32Div(x, y, dst)
	}
}

func BenchmarkVec8x32AddPureGo(b *testing.B) {
	x := &[8]float32{1, 2, 3, 4, 5, 6, 7, 8}
	y := &[8]float32{5, 6, 7, 8, 9, 5, 6, 8}
	dst := &[8]float32{}
	for b.Loop() {
		for i := range 8 {
			dst[i] = x[i] + y[i]
		}
	}
}

func BenchmarkVec8x32SubPureGo(b *testing.B) {
	x := &[8]float32{1, 2, 3, 4, 5, 6, 7, 8}
	y := &[8]float32{5, 6, 7, 8, 9, 5, 6, 8}
	dst := &[8]float32{}
	for b.Loop() {
		for i := range 8 {
			dst[i] = x[i] - y[i]
		}
	}
}

func BenchmarkVec8x32MulPureGo(b *testing.B) {
	x := &[8]float32{1, 2, 3, 4, 5, 6, 7, 8}
	y := &[8]float32{5, 6, 7, 8, 9, 5, 6, 8}
	dst := &[8]float32{}
	for b.Loop() {
		for i := range 8 {
			dst[i] = x[i] * y[i]
		}
	}
}

func BenchmarkVec8x32DivPureGo(b *testing.B) {
	x := &[8]float32{1, 2, 3, 4, 5, 6, 7, 8}
	y := &[8]float32{5, 6, 7, 8, 9, 5, 6, 8}
	dst := &[8]float32{}
	for b.Loop() {
		for i := range 8 {
			dst[i] = x[i] / y[i]
		}
	}
}
