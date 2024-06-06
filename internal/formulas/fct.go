package formulas

type Distributor float64

const (
	LISBOAGAS    Distributor = 17
	SETGAS       Distributor = 16
	LUSITANIAGAS Distributor = 16
	MEDIGAS      Distributor = 18
	PAXGAS       Distributor = 15
	DIANAGAS     Distributor = 16
	DURIENGAS    Distributor = 11
	BEIRAGAS     Distributor = 13
)

const temperatureRatio = 273.15

func CalculateTemperatureCorrection(temperature Distributor) float64 {
	return float64(temperatureRatio / (temperatureRatio + temperature))
}
