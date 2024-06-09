package ren

import (
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/domain/entity"
)

type RenReadingRepository struct {
}

func NewRenReadingRepository() RenReadingRepository {
	return RenReadingRepository{}
}

func (r RenReadingRepository) FindByInterval(
	interval entity.Interval,
) []entity.Reading {
	csvReader, err := downloadCsv(interval)
	output := []entity.Reading{}
	if err != nil {
		return output
	}
	defer csvReader.Close()
	parsedReadings, err := parseCsv(csvReader)
	if err != nil {
		return output
	}
	return parsedReadings
}
