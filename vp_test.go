package vp

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		List   []float64
		Expect float64
	}{
		{[]float64{0, 0, 0}, 0},
		{[]float64{1, 0, 0}, 1},
		{[]float64{1, 2, 3}, 3},
		{[]float64{1, 3, 2}, 3},
	}
	for _, tt := range tests {
		got := Max(tt.List)
		if got != tt.Expect {
			t.Errorf("Max %v expect %f, got %f", tt.List, tt.Expect, got)
		}
	}
}
func TestMin(t *testing.T) {
	tests := []struct {
		List   []float64
		Expect float64
	}{
		{[]float64{0, 0, 0}, 0},
		{[]float64{1, 0, 0}, 0},
		{[]float64{1, 2, 3}, 1},
		{[]float64{1, 3, 2}, 1},
	}
	for _, tt := range tests {
		got := Min(tt.List)
		if got != tt.Expect {
			t.Errorf("Max %v expect %f, got %f", tt.List, tt.Expect, got)
		}
	}
}
