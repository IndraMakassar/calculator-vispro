package mathops

import "testing"

func TestAdd(t *testing.T) {
	got := Add(4, 6)
	want := 10.0
	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(10, 12)
	want := -2.0
	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(8, 5)
	want := 40.0
	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}

func TestDivide(t *testing.T) {
	got, _ := Divide(10, 5)
	want := 2.0
	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}

func TestDivide2(t *testing.T) {
	got, errors := Divide(10, 0)
	want := 0.0
	if got != want || errors == nil {
		t.Errorf("got %g, wanted %g", got, want)
	}
}
