package generation

import "troyan/pkg/dtypes"

func GetCorrelatedSignal() (res []float64) {
	res = make([]float64, 0, dtypes.Dots)
	for i := 0; i < dtypes.Dots; i++ {
		res = append(res, float64(i))
	}
	return res
}
