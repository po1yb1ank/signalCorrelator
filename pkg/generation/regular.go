package generation

import (
	"math"
	"troyan/pkg/dtypes"
)

func GetRegularSignal() []float64 {
	vals := make([]float64, dtypes.Dots)
	i := 1.0
	for k := range vals {
		vals[k] = math.Sin(i)
		i++
	}
	return vals
}
