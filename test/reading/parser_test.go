package reading

import (
	"slices"
	"testing"
	"time"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/reading"
)

func TestPruneExcessValues(t *testing.T) {
	lower, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}

	startTime := lower.AddDate(0, 0, 1)
	middleTime := lower.AddDate(0, 0, 2)
	stopTime := lower.AddDate(0, 0, 3)
	higher := lower.AddDate(0, 0, 4)

	input := []reading.Reading{
		reading.NewReading(lower, "CTS 6000 Valença do Minho", 0),
		reading.NewReading(lower, "CTS 6000 Valença do Minho", 0),
		reading.NewReading(startTime, "CTS 6000 Valença do Minho", 1),
		reading.NewReading(middleTime, "CTS 6000 Valença do Minho", 2),
		reading.NewReading(stopTime, "CTS 6000 Valença do Minho", 3),
		reading.NewReading(higher, "CTS 6000 Valença do Minho", 4),
		reading.NewReading(higher, "CTS 6000 Valença do Minho", 4),
	}

	expected := slices.DeleteFunc(input, func(r reading.Reading) bool {
		for _, time := range []time.Time{startTime, middleTime, stopTime} {
			if r.CompareDay(time) == 0 {
				return false
			}
		}
		return true
	})

	actual := reading.PruneExcessValues(input, startTime, stopTime)
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

	input := []reading.Reading{}

	expected := []reading.Reading{}
	actual := reading.PruneExcessValues(input, startTime, stopTime)
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

	input := []reading.Reading{
		reading.NewReading(lower, "CTS 6000 Valença do Minho", 0),
		reading.NewReading(lower, "CTS 6000 Valença do Minho", 0),
		reading.NewReading(higher, "CTS 6000 Valença do Minho", 4),
		reading.NewReading(higher, "CTS 6000 Valença do Minho", 4),
	}

	expected := []reading.Reading{}

	actual := reading.PruneExcessValues(input, startTime, stopTime)
	if len(expected) != len(actual) {
		t.Errorf("expected %d, actual %d", len(expected), len(actual))
	}
}
