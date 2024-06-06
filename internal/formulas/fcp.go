package formulas

const pressureRatio = 1.01325

func CalculatePressureCorrectionBar(pr float64) float64 {
	return (pr + pressureRatio) / pressureRatio
}

func CalculatePressureCorrectionMilliBar(pr float64) float64 {
	return CalculatePressureCorrectionBar(pr / 1000)
}
