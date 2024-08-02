package ren

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/domain/entity"
)

const timelayout = "2006-01-02 15:04"

func (r RenReadingRepository) parseCsv(
	reader io.Reader,
) ([]entity.Reading, error) {
	lines := []entity.Reading{}
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

		lines = append(
			lines,
			entity.NewReading(timestamp, splittedLines[1], number),
		)
	}
	return lines, nil
}
