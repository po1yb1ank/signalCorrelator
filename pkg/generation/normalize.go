package generation

import (
	"log"
	"troyan/pkg/dtypes"
)

func Normalize(signal []float64) []float64 {
	var avgX float64
	for _, v := range signal {
		avgX += float64(v)
	}
	avgX = float64(1) / dtypes.Dots * avgX
	log.Println("AvgX", avgX)
	centratedX := make([]float64, 0, len(signal))
	for _, v := range signal {
		centratedX = append(centratedX, float64(v)-avgX)
	}
	log.Println("centratedX", centratedX)
	rxx := make([]float64, 0, len(signal))
	var stepSum float64

	for i := 0; i < cap(rxx); i++ {
		for j := 1; j < dtypes.Dots-i; j++ {
			stepSum += centratedX[j] * centratedX[i+j]
		}
		rxx = append(rxx, (1.0/float64(dtypes.Dots-i))*stepSum)
	}
	return rxx
}
