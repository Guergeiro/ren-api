package repository

import (
	"context"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/domain/entity"
)

type ReadingRepository interface {
	FindByInterval(
		ctx context.Context,
		interval entity.Interval,
	) []entity.Reading
}
