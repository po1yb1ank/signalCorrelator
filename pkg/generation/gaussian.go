package generation

import (
	"math/rand"
	"troyan/pkg/dtypes"
)

func GetGaussianSignal() []float64 {
	var res []float64
	for i := 0; i < dtypes.Dots; i++ {
		res = append(res, rand.NormFloat64())
	}
	return res
}
