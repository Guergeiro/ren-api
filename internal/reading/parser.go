package reading

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
	"time"
)

const timelayout = "2006-01-02 15:04"

func ParseCsv(reader io.Reader) ([]Reading, error) {
	lines := []Reading{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		splittedLines := strings.Split(scanner.Text(), ";")
		if splittedLines[2] != "PCS kWh/m3" {
			continue
		}
		stdtime := strings.ReplaceAll(splittedLines[0], "/", "-")
		timestamp, err := time.Parse(timelayout, stdtime)
		if err != nil {
			return nil, err
		}

		stdstr := strings.ReplaceAll(splittedLines[3], ",", ".")
		number, err := strconv.ParseFloat(stdstr, 64)
		if err != nil {
			return nil, err
		}

		lines = append(lines, NewReading(timestamp, splittedLines[1], number))
	}
	return lines, nil
}

func PruneExcessValues(
	readings []Reading,
	startTime,
	stopTime time.Time,
) []Reading {
	return slices.DeleteFunc(readings, func(reading Reading) bool {
		if reading.CompareDay(startTime) < 0 {
			return true
		}
		if reading.CompareDay(stopTime) > 0 {
			return true
		}
		return false
	})
}
