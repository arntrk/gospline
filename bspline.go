package gospline

import (
	"fmt"

	bs "github.com/arntrk/gobasis"
)

type BSpline struct {
	pnts [][]float64

	basis *bs.BSplineBasis
}

// NewBSpline create new BSplineCurve
func NewBSpline(pnts [][]float64, order int) (*BSpline, error) {

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

	// fmt.Printf("knts: %v\n", knts)
	// fmt.Println("size: ", len(knts), " cap: ", cap(knts))

	b, err := bs.Create(knts, order)
	if err != nil {
		return nil, err
	}

	return &BSpline{pnts: pnts, basis: b}, nil
}

func (bs *BSpline) Eval(t float64) []float64 {

	i, knts := bs.basis.Eval(t)
	res := make([]float64, len(bs.pnts[0]))

	for k := 0; k < len(knts); k++ {
		for j := 0; j < len(bs.pnts[i+k]); j++ {
			res[j] += bs.pnts[i+k][j] * knts[k]
		}
	}

	return res
}

// Interval bla bla
func (bs *BSpline) Interval() (float64, float64) {

	return bs.basis.Interval()
}

// Derivate bla bla
func (bs *BSpline) Derivate() *BSpline {

	basis := bs.basis.Derivate()

	if basis != nil {

		pnts := make([][]float64, len(bs.pnts)-1)

		for i := 0; i < len(pnts); i++ {
			pnts[i] = make([]float64, len(bs.pnts[i]))

			p := bs.basis.Order() - 1
			k := float64(p) / (bs.basis.Knot(i+p+1) - bs.basis.Knot(i+1))

			// assumes 3d coordinates for now
			pnts[i][0] = k * (bs.pnts[i+1][0] - bs.pnts[i][0])
			pnts[i][1] = k * (bs.pnts[i+1][1] - bs.pnts[i][1])
			pnts[i][2] = k * (bs.pnts[i+1][2] - bs.pnts[i][2])

		}

		b := &BSpline{basis: basis, pnts: pnts}

		return b
	}

	return nil
}
