package pcs

import (
	"cmp"
	"context"
	"testing"
	"time"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/application/service"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/application/usecase/pcs"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/domain/entity"
)

type mock_repository struct {
	returnValue []entity.Reading
}

func (r mock_repository) FindByInterval(
	interval entity.Interval,
) []entity.Reading {
	return r.returnValue
}

const timelayout = "2006-01-02"

func TestReturnsZeroAverage(t *testing.T) {
	time, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}

	usecase := pcs.NewGetAverageUseCase(
		mock_repository{
			returnValue: []entity.Reading{},
		},
		service.NewReadingPruner(),
	)

	expected := 0.0

	actual, err := usecase.Execute(
		pcs.NewAverageProps(context.Background(), time, time.AddDate(0, 1, 0)),
	)
	if err != nil {
		t.Fail()
	}
	if cmp.Compare(expected, actual) != 0 {
		t.Errorf("expected %f, actual %f", expected, actual)
	}
}

func TestReturnsFiveAverage(t *testing.T) {
	time, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}

	usecase := pcs.NewGetAverageUseCase(
		mock_repository{
			returnValue: []entity.Reading{
				entity.NewReading(time, string(entity.CAMPO_MAIOR), 10.0),
				entity.NewReading(time, string(entity.CAMPO_MAIOR), 0.0),
			},
		},
		service.NewReadingPruner(),
	)

	expected := 5.0

	actual, err := usecase.Execute(
		pcs.NewAverageProps(context.Background(), time, time.AddDate(0, 1, 0)),
	)
	if err != nil {
		t.Fail()
	}
	if cmp.Compare(expected, actual) != 0 {
		t.Errorf("expected %f, actual %f", expected, actual)
	}
}
