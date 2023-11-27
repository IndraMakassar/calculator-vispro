package calculator

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
	got := Divide(10, 5)
	want := 2.0
	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}

func TestDivide2(t *testing.T) {
	got := Divide(10, 0)
	want := 0.0
	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}

func TestPower(t *testing.T) {
	got := Power(5, 2)
	want := 25.0
	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}

func TestPower2(t *testing.T) {
	got := Power(5, 0)
	want := 1.0
	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}

func TestPower3(t *testing.T) {
	got := Power(5, 1)
	want := 5.0
	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}

func TestPower4(t *testing.T) {
	got := Power(5, 3)
	want := 125.0
	if got != want {
		t.Errorf("got %g, wanted %g", got, want)
	}
}
