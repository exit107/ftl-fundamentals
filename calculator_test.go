package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
		a, b float64
		want float64
	}

	testCases := []testCase{
		{name: "test case name", a: 2, b: 2, want: 4},
		{name: "test case name", a: 1, b: 5, want: 6},
		{name: "test case name", a: 5, b: 0, want: 5},
		{name: "test case name", a: -5, b: 1, want: -4},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}

	testCases := []testCase{
		{name: "test case name", a: 2, b: 2, want: 0},
		{name: "test case name", a: 1, b: 5, want: -4},
		{name: "test case name", a: 5, b: 0, want: 5},
		{name: "test case name", a: -5, b: 1, want: -6},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}

	testCases := []testCase{
		{name: "test case name", a: 2, b: 2, want: 4},
		{name: "test case name", a: 1, b: 5, want: 5},
		{name: "test case name", a: 5, b: 0, want: 0},
		{name: "test case name", a: -5, b: 1, want: -5},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name        string
		a, b        float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{name: "test case name", a: 2, b: 2, want: 1, errExpected: false},
		{name: "test case name", a: 1, b: 5, want: .2, errExpected: false},
		{name: "test case name", a: 5, b: 0, want: 0, errExpected: true},
		{name: "test case name", a: -5, b: 1, want: -5, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil && !tc.errExpected {
			t.Errorf("Multiply(%f, %f): Unexpected error returned.", tc.a, tc.b)
		}
		if err == nil && tc.errExpected {
			t.Errorf("Multiply(%f, %f): Expected error not found.", tc.a, tc.b)
		}
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
