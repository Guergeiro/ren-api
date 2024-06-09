package service

import (
	"slices"
	"time"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/domain/entity"
)

type ReadingPruner struct {
}

func NewReadingPruner() ReadingPruner {
	return ReadingPruner{}
}

func (p ReadingPruner) PruneExcessValues(
	readings []entity.Reading,
	startTime,
	stopTime time.Time,
) []entity.Reading {
	return slices.DeleteFunc(readings, func(reading entity.Reading) bool {
		if reading.CompareDay(startTime) < 0 {
			return true
		}
		if reading.CompareDay(stopTime) > 0 {
			return true
		}
		return false
	})
}
