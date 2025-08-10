package simdgo_test

import (
	"reflect"
	"testing"

	"github.com/xtt28/simdgo"
)

func TestVec4x32Add(t *testing.T) {
	tests := []struct {
		a        [4]float32
		b        [4]float32
		expected [4]float32
	}{
		{
			[...]float32{1, 2, 3, 4},
			[...]float32{5, 6, 7, 8},
			[...]float32{6, 8, 10, 12},
		},
		{
			[...]float32{5.1, 2.3, 9.6, 8.5},
			[...]float32{2.9, 8.4, 0.2, 3.1},
			[...]float32{8.0, 10.7, 9.8, 11.6},
		},
	}

	for _, test := range tests {
		dst := &[4]float32{}
		simdgo.Vec4x32Add(&test.a, &test.b, dst)

		if !reflect.DeepEqual(*dst, test.expected) {
			t.Fatalf("Vec4x32Add(%v, %v): expected %v, got %v", test.a, test.b, test.expected, *dst)
		}
	}
}

func TestVec4x32Sub(t *testing.T) {
	tests := []struct {
		a        [4]float32
		b        [4]float32
		expected [4]float32
	}{
		{
			[...]float32{5.4, 6, 7, 8},
			[...]float32{1, 5, 2, 9},
			[...]float32{4.4, 1, 5, -1},
		},
	}

	for _, test := range tests {
		dst := &[4]float32{}
		simdgo.Vec4x32Sub(&test.a, &test.b, dst)

		if !reflect.DeepEqual(*dst, test.expected) {
			t.Fatalf("Vec4x32Sub(%v, %v): expected %v, got %v", test.a, test.b, test.expected, *dst)
		}
	}
}

func TestVec4x32Mul(t *testing.T) {
	tests := []struct {
		a        [4]float32
		b        [4]float32
		expected [4]float32
	}{
		{
			[...]float32{5.4, 6, 7, 8},
			[...]float32{1, 5, 2, 9},
			[...]float32{5.4, 30, 14, 72},
		},
	}

	for _, test := range tests {
		dst := &[4]float32{}
		simdgo.Vec4x32Mul(&test.a, &test.b, dst)

		if !reflect.DeepEqual(*dst, test.expected) {
			t.Fatalf("Vec4x32Mul(%v, %v): expected %v, got %v", test.a, test.b, test.expected, *dst)
		}
	}
}

func TestVec4x32Div(t *testing.T) {
	tests := []struct {
		a        [4]float32
		b        [4]float32
		expected [4]float32
	}{
		{
			[...]float32{5.4, 6, 7, 8},
			[...]float32{2, 3, 2, 5},
			[...]float32{2.7, 2, 3.5, 1.6},
		},
	}

	for _, test := range tests {
		dst := &[4]float32{}
		simdgo.Vec4x32Div(&test.a, &test.b, dst)

		if !reflect.DeepEqual(*dst, test.expected) {
			t.Fatalf("Vec4x32Div(%v, %v): expected %v, got %v", test.a, test.b, test.expected, *dst)
		}
	}
}

func BenchmarkVec4x32Add(b *testing.B) {
	x := &[4]float32{1, 2, 3, 4}
	y := &[4]float32{5, 6, 7, 8}
	dst := &[4]float32{}
	for b.Loop() {
		simdgo.Vec4x32Add(x, y, dst)
	}
}

func BenchmarkVec4x32Sub(b *testing.B) {
	x := &[4]float32{1, 2, 3, 4}
	y := &[4]float32{5, 6, 7, 8}
	dst := &[4]float32{}
	for b.Loop() {
		simdgo.Vec4x32Sub(x, y, dst)
	}
}

func BenchmarkVec4x32Mul(b *testing.B) {
	x := &[4]float32{1, 2, 3, 4}
	y := &[4]float32{5, 6, 7, 8}
	dst := &[4]float32{}
	for b.Loop() {
		simdgo.Vec4x32Mul(x, y, dst)
	}
}

func BenchmarkVec4x32Div(b *testing.B) {
	x := &[4]float32{1, 2, 3, 4}
	y := &[4]float32{5, 6, 7, 8}
	dst := &[4]float32{}
	for b.Loop() {
		simdgo.Vec4x32Div(x, y, dst)
	}
}

func BenchmarkVec4x32AddPureGo(b *testing.B) {
	x := &[4]float32{1, 2, 3, 4}
	y := &[4]float32{5, 6, 7, 8}
	dst := &[4]float32{}
	for b.Loop() {
		for i := range 4 {
			dst[i] = x[i] + y[i]
		}
	}
}

func BenchmarkVec4x32SubPureGo(b *testing.B) {
	x := &[4]float32{1, 2, 3, 4}
	y := &[4]float32{5, 6, 7, 8}
	dst := &[4]float32{}
	for b.Loop() {
		for i := range 4 {
			dst[i] = x[i] - y[i]
		}
	}
}

func BenchmarkVec4x32MulPureGo(b *testing.B) {
	x := &[4]float32{1, 2, 3, 4}
	y := &[4]float32{5, 6, 7, 8}
	dst := &[4]float32{}
	for b.Loop() {
		for i := range 4 {
			dst[i] = x[i] * y[i]
		}
	}
}

func BenchmarkVec4x32DivPureGo(b *testing.B) {
	x := &[4]float32{1, 2, 3, 4}
	y := &[4]float32{5, 6, 7, 8}
	dst := &[4]float32{}
	for b.Loop() {
		for i := range 4 {
			dst[i] = x[i] / y[i]
		}
	}
}
