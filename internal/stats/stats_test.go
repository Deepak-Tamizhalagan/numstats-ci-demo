package stats

import "testing"

func TestSum(t *testing.T) {
	got := Sum([]float64{1, 2, 3})
	want := 6.0
	if got != want {
		t.Fatalf("Sum() = %v; want %v", got, want)
	}
}
func TestMean(t *testing.T) {
	got := Mean([]float64{2, 4, 6})
	want := 4.0
	if got != want {
		t.Fatalf("Mean() = %v; want %v", got, want)
	}
}
func TestMedianOdd(t *testing.T) {
	got := Median([]float64{3, 1, 2})
	want := 2.0
	if got != want {
		t.Fatalf("Median(odd) = %v; want %v", got, want)
	}
}

func TestMedianEven(t *testing.T) {
	got := Median([]float64{1, 2, 3, 4})
	want := 2.5
	if got != want {
		t.Fatalf("Median(even) = %v; want %v", got, want)
	}
}
func TestMinMax(t *testing.T) {
	nums := []float64{5, 1, 8, -2, 7}
	if got := Min(nums); got != -2 {
		t.Fatalf("Min() = %v; want -2", got)
	}
	if got := Max(nums); got != 8 {
		t.Fatalf("Max() = %v; want 8", got)
	}
}
