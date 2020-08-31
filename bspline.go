package gospline

import (
	"fmt"

	bs "github.com/arntrk/gobasis"
)

type bspline struct {
	pnts [][]float64

	basis *bs.BSplineBasis
}

// NewBSpline create new BSplineCurve
func NewBSpline(pnts [][]float64, order int) (*bspline, error) {

	if len(pnts) < order {
		return nil, fmt.Errorf("Minimum number of points must match the order (%d >= %d)", len(pnts), order)
	}

	// 2*order + (len(pnts) - order) num elements
	//knts := make([]float64, len(pnts) + order)

	knts := make([]float64, 0, len(pnts)+order)

	// append start knts
	for i := 0; i < order; i++ {
		knts = append(knts, 0.0)
	}

	steps := (1.0 - 0.0) / float64(len(pnts)-order+1)

	for i := 1; i <= (len(pnts) - order); i++ {
		knts = append(knts, float64(i)*steps)
	}

	// append stop knts
	for i := 0; i < order; i++ {
		knts = append(knts, 1.0)
	}

	fmt.Printf("knts: %v\n", knts)
	fmt.Println("size: ", len(knts), " cap: ", cap(knts))

	b, err := bs.Create(knts, order)
	if err != nil {
		return nil, err
	}

	return &bspline{pnts: pnts, basis: b}, nil
}

func (bs *bspline) Eval(t float64) []float64 {

	i, knts := bs.basis.Eval(t)
	res := make([]float64, len(bs.pnts[0]))

	for k := 0; k < len(knts); k++ {
		for j := 0; j < len(bs.pnts[i+k]); j++ {
			res[j] += bs.pnts[i+k][j] * knts[k]
		}
	}

	return res
}

func (bs *bspline) Interval() (float64, float64) {

	return bs.basis.Interval()
}
