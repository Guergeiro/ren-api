package pcs

import (
	"sync"
	"time"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/interval"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/reading"
)

func Average(startTime, stopTime time.Time) (float64, error) {
	intervals := interval.NewIntervals(
		startTime.AddDate(0, 0, -1),
		stopTime.AddDate(0, 0, 1),
	)
	channel := make(chan reading.Reading)
	wg := sync.WaitGroup{}
	for _, in := range intervals {
		wg.Add(1)
		go func(in interval.Interval, wg *sync.WaitGroup) {
			defer wg.Done()
			csvReader, err := reading.DownloadCsv(in)
			if err != nil {
				return
			}
			defer csvReader.Close()
			parsedReadings, err := reading.ParseCsv(csvReader)
			if err != nil {
				return
			}
			for _, reading := range reading.PruneExcessValues(
				parsedReadings,
				startTime,
				stopTime,
			) {
				channel <- reading
			}
		}(in, &wg)
	}
	go func() {
		wg.Wait()
		close(channel)
	}()
	count := 0
	total := 0.0
	for reading := range channel {
		total += reading.Value()
		count += 1
	}
	average := total / float64(count)
	return average, nil
}
