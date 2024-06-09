package service

import (
	"slices"
	"testing"
	"time"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/application/service"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/domain/entity"
)

const timelayout = "2006-01-02"

func TestPruneExcessValues(t *testing.T) {
	lower, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}

	startTime := lower.AddDate(0, 0, 1)
	middleTime := lower.AddDate(0, 0, 2)
	stopTime := lower.AddDate(0, 0, 3)
	higher := lower.AddDate(0, 0, 4)

	input := []entity.Reading{
		entity.NewReading(lower, "CTS 6000 Valença do Minho", 0),
		entity.NewReading(lower, "CTS 6000 Valença do Minho", 0),
		entity.NewReading(startTime, "CTS 6000 Valença do Minho", 1),
		entity.NewReading(middleTime, "CTS 6000 Valença do Minho", 2),
		entity.NewReading(stopTime, "CTS 6000 Valença do Minho", 3),
		entity.NewReading(higher, "CTS 6000 Valença do Minho", 4),
		entity.NewReading(higher, "CTS 6000 Valença do Minho", 4),
	}

	expected := slices.DeleteFunc(input, func(r entity.Reading) bool {
		for _, time := range []time.Time{startTime, middleTime, stopTime} {
			if r.CompareDay(time) == 0 {
				return false
			}
		}
		return true
	})

	actual := service.NewReadingPruner().PruneExcessValues(input, startTime, stopTime)
	if len(expected) != len(actual) {
		t.Errorf("expected %d, actual %d", len(expected), len(actual))
	}
}

func TestPruneExcessValuesEmpty(t *testing.T) {
	lower, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}

	startTime := lower.AddDate(0, 0, 1)
	stopTime := lower.AddDate(0, 0, 3)

	input := []entity.Reading{}

	expected := []entity.Reading{}
	actual := service.NewReadingPruner().PruneExcessValues(input, startTime, stopTime)
	if len(expected) != len(actual) {
		t.Errorf("expected %d, actual %d", len(expected), len(actual))
	}
}

func TestPruneExcessValuesNoValuesInRange(t *testing.T) {
	lower, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}

	startTime := lower.AddDate(0, 0, 1)
	stopTime := lower.AddDate(0, 0, 3)
	higher := lower.AddDate(0, 0, 4)

	input := []entity.Reading{
		entity.NewReading(lower, "CTS 6000 Valença do Minho", 0),
		entity.NewReading(lower, "CTS 6000 Valença do Minho", 0),
		entity.NewReading(higher, "CTS 6000 Valença do Minho", 4),
		entity.NewReading(higher, "CTS 6000 Valença do Minho", 4),
	}

	expected := []entity.Reading{}

	actual := service.NewReadingPruner().PruneExcessValues(input, startTime, stopTime)
	if len(expected) != len(actual) {
		t.Errorf("expected %d, actual %d", len(expected), len(actual))
	}
}
