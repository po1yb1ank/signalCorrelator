package generation

import (
	"math/rand"
	"troyan/pkg/dtypes"
)

func GetExampleSignal() []float64 {
	generated := make([]float64, dtypes.Dots)

	rand.Seed(rand.Int63())
	for i := 0; i < dtypes.Dots; i++ {
		generated[i] = float64(rand.Intn(2))
	}
	return generated
}
