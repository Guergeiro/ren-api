package ren

import (
	"context"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/domain/entity"
)

type RenReadingRepository struct {
	endpoint string
}

func NewRenReadingRepository(endpoint string) RenReadingRepository {
	return RenReadingRepository{
		endpoint,
	}
}

func (r RenReadingRepository) FindByInterval(
	ctx context.Context,
	interval entity.Interval,
) []entity.Reading {
	csvReader, err := r.downloadCsv(interval)
	output := []entity.Reading{}
	if err != nil {
		return output
	}
	defer csvReader.Close()
	parsedReadings, err := r.parseCsv(csvReader)
	if err != nil {
		return output
	}
	return parsedReadings
}
