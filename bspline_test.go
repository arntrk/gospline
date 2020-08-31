package gospline

import (
	"fmt"
	"testing"
)

func TestBSpline(t *testing.T) {

	pnts := [][]float64{
		{0.0, 100.0, -1000.0},
		{75.0, 200.0, -750.0},
		{-75.0, 150.0, -500.0},
		{40.0, 75.0, -100.0},
		{-40.0, 30.0, -50.0},
		{5.0, 30.0, 20.0},
	}

	spline, err := NewBSpline(pnts, 4)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	var tt float64 = 0.0

	for tt <= 1.0 {
		p := spline.Eval(tt)
		fmt.Printf("p := %v at t: %f\n", p, tt)
		tt += 0.1
	}

	if err == nil {
		t.Errorf("Error: %v", err)
	}
}
