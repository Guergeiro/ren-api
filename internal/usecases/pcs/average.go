package pcs

import (
	"time"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/reading"
)

func Average(startTime, stopTime time.Time) (float64, error) {
	csvReader, err := reading.DownloadCsv(startTime, stopTime)
	if err != nil {
		return 0, err
	}
	defer csvReader.Close()
	readings, err := reading.ParseCsv(csvReader)
	if err != nil {
		return 0, err
	}
	readings = reading.PruneExcessValues(readings, startTime, stopTime)
	total := 0.0
	for _, reading := range readings {
		total += reading.Value()
	}
	average := total / float64(len(readings))
	return average, nil
}
