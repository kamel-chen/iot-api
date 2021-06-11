package num

import (
	"math/rand"
	"time"
)

func GPSRound(num float64) float64 {
	return float64(int(num * 10000000)) / 10000000
}

func RandomFloat(max float64, min float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64() * (max - min) + min
}
