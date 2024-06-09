package repository

import "github.com/guergeiro/fator-conversao-gas-portugal/internal/domain/entity"

type ReadingRepository interface {
	FindByInterval(interval entity.Interval) []entity.Reading
}
