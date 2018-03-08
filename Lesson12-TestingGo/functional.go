package functional

func Average(xs []float64) float64 {
	var summer, counter float64
	for _, v := range xs {
		summer = summer + v
		counter++
	}
	return (summer / counter)
}
