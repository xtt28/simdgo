package simdgo_test

import (
	"reflect"
	"testing"

	"github.com/xtt28/simdgo"
)

func TestVec4x64Add(t *testing.T) {
	tests := []struct {
		a        [4]float64
		b        [4]float64
		expected [4]float64
	}{
		{
			[...]float64{1, 2, 3, 4},
			[...]float64{5, 6, 7, 8},
			[...]float64{6, 8, 10, 12},
		},
		{
			[...]float64{5.1, 2.3, 9.6, 8.5},
			[...]float64{2.9, 8.4, 0.1, 3.1},
			[...]float64{8.0, 10.7, 9.7, 11.6},
		},
	}

	for _, test := range tests {
		dst := &[4]float64{}
		simdgo.Vec4x64Add(&test.a, &test.b, dst)

		if !reflect.DeepEqual(*dst, test.expected) {
			t.Fatalf("Vec4x64Add(%v, %v): expected %v, got %v", test.a, test.b, test.expected, *dst)
		}
	}
}

func TestVec4x64Sub(t *testing.T) {
	tests := []struct {
		a        [4]float64
		b        [4]float64
		expected [4]float64
	}{
		{
			[...]float64{5.4, 6, 7, 8},
			[...]float64{1, 5, 2, 9},
			[...]float64{4.4, 1, 5, -1},
		},
	}

	for _, test := range tests {
		dst := &[4]float64{}
		simdgo.Vec4x64Sub(&test.a, &test.b, dst)

		if !reflect.DeepEqual(*dst, test.expected) {
			t.Fatalf("Vec4x64Sub(%v, %v): expected %v, got %v", test.a, test.b, test.expected, *dst)
		}
	}
}

func TestVec4x64Mul(t *testing.T) {
	tests := []struct {
		a        [4]float64
		b        [4]float64
		expected [4]float64
	}{
		{
			[...]float64{5.4, 6, 7, 8},
			[...]float64{1, 5, 2, 9},
			[...]float64{5.4, 30, 14, 72},
		},
	}

	for _, test := range tests {
		dst := &[4]float64{}
		simdgo.Vec4x64Mul(&test.a, &test.b, dst)

		if !reflect.DeepEqual(*dst, test.expected) {
			t.Fatalf("Vec4x64Mul(%v, %v): expected %v, got %v", test.a, test.b, test.expected, *dst)
		}
	}
}

func TestVec4x64Div(t *testing.T) {
	tests := []struct {
		a        [4]float64
		b        [4]float64
		expected [4]float64
	}{
		{
			[...]float64{5.4, 6, 7, 8},
			[...]float64{2, 3, 2, 5},
			[...]float64{2.7, 2, 3.5, 1.6},
		},
	}

	for _, test := range tests {
		dst := &[4]float64{}
		simdgo.Vec4x64Div(&test.a, &test.b, dst)

		if !reflect.DeepEqual(*dst, test.expected) {
			t.Fatalf("Vec4x64Div(%v, %v): expected %v, got %v", test.a, test.b, test.expected, *dst)
		}
	}
}

func BenchmarkVec4x64Add(b *testing.B) {
	x := &[4]float64{1, 2, 3, 4}
	y := &[4]float64{5, 6, 7, 8}
	dst := &[4]float64{}
	for b.Loop() {
		simdgo.Vec4x64Add(x, y, dst)
	}
}

func BenchmarkVec4x64Sub(b *testing.B) {
	x := &[4]float64{1, 2, 3, 4}
	y := &[4]float64{5, 6, 7, 8}
	dst := &[4]float64{}
	for b.Loop() {
		simdgo.Vec4x64Sub(x, y, dst)
	}
}

func BenchmarkVec4x64Mul(b *testing.B) {
	x := &[4]float64{1, 2, 3, 4}
	y := &[4]float64{5, 6, 7, 8}
	dst := &[4]float64{}
	for b.Loop() {
		simdgo.Vec4x64Mul(x, y, dst)
	}
}

func BenchmarkVec4x64Div(b *testing.B) {
	x := &[4]float64{1, 2, 3, 4}
	y := &[4]float64{5, 6, 7, 8}
	dst := &[4]float64{}
	for b.Loop() {
		simdgo.Vec4x64Div(x, y, dst)
	}
}

func BenchmarkVec4x64AddPureGo(b *testing.B) {
	x := &[4]float64{1, 2, 3, 4}
	y := &[4]float64{5, 6, 7, 8}
	dst := &[4]float64{}
	for b.Loop() {
		for i := range 4 {
			dst[i] = x[i] + y[i]
		}
	}
}

func BenchmarkVec4x64SubPureGo(b *testing.B) {
	x := &[4]float64{1, 2, 3, 4}
	y := &[4]float64{5, 6, 7, 8}
	dst := &[4]float64{}
	for b.Loop() {
		for i := range 4 {
			dst[i] = x[i] - y[i]
		}
	}
}

func BenchmarkVec4x64MulPureGo(b *testing.B) {
	x := &[4]float64{1, 2, 3, 4}
	y := &[4]float64{5, 6, 7, 8}
	dst := &[4]float64{}
	for b.Loop() {
		for i := range 4 {
			dst[i] = x[i] * y[i]
		}
	}
}

func BenchmarkVec4x64DivPureGo(b *testing.B) {
	x := &[4]float64{1, 2, 3, 4}
	y := &[4]float64{5, 6, 7, 8}
	dst := &[4]float64{}
	for b.Loop() {
		for i := range 4 {
			dst[i] = x[i] / y[i]
		}
	}
}
