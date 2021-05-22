package helpers

func IsRateValid(rate float64) bool {
	return rate >= 0.9 && rate <= 1.1
}
