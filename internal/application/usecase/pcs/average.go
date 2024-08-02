package pcs

import (
	"context"
	"time"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/application/service"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/domain/entity"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/domain/repository"

	"golang.org/x/sync/errgroup"
)

type AverageProps struct {
	ctx       context.Context
	startTime time.Time
	stopTime  time.Time
}

func NewAverageProps(
	ctx context.Context,
	startTime,
	stopTime time.Time,
) AverageProps {
	return AverageProps{
		ctx,
		startTime,
		stopTime,
	}
}

type GetAverageUseCase struct {
	repository repository.ReadingRepository
	pruner     service.ReadingPruner
}

func NewGetAverageUseCase(
	repository repository.ReadingRepository,
	pruner service.ReadingPruner,
) GetAverageUseCase {
	return GetAverageUseCase{
		repository,
		pruner,
	}
}

func (u GetAverageUseCase) Execute(props AverageProps) (float64, error) {
	intervals := entity.NewIntervals(
		props.startTime.AddDate(0, 0, -1),
		props.stopTime.AddDate(0, 0, 1),
	)
	channel := make(chan entity.Reading)
	wg, ctx := errgroup.WithContext(props.ctx)
	for _, in := range intervals {
		in := in
		wg.Go(func() error {
			parsedReadings := u.repository.FindByInterval(
				ctx,
				in,
			)
			for _, reading := range u.pruner.PruneExcessValues(
				parsedReadings,
				props.startTime,
				props.stopTime,
			) {
				channel <- reading
			}
			return nil
		})
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
	if count == 0 {
		return total, nil
	}
	average := total / float64(count)
	return average, nil
}
