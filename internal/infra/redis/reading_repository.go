package redis

import (
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/domain/entity"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/domain/repository"
)

type RedisReadingRepository struct {
	repository repository.ReadingRepository
}

func NewRedisReadingRepository(
	repository repository.ReadingRepository,
) RedisReadingRepository {
	return RedisReadingRepository{
		repository,
	}
}

func (r RedisReadingRepository) FindByInterval(
	interval entity.Interval,
) []entity.Reading {
	return r.repository.FindByInterval(interval)
}
