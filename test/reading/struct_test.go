package reading

import (
	"testing"
	"time"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/reading"
)

const timelayout = "2006-01-02"

func TestEqualReadings(t *testing.T) {
	time, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}
	r1 := reading.NewReading(time, "CTS 6000 Valença do Minho", 0)
	r2 := reading.NewReading(time, "CTS 6000 Valença do Minho", 0)

	expected := true
	actual := r1.Equal(r2)

	if expected != actual {
		t.Errorf("expected %t, actual %t", expected, actual)
	}
}
func TestEqualDifferentReadings(t *testing.T) {
	time, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}
	base := reading.NewReading(time, "CTS 6000 Valença do Minho", 0)

	input := []reading.Reading{
		reading.NewReading(
			time.AddDate(0, 0, 1), "CTS 6000 Valença do Minho", 0,
		),
		reading.NewReading(
			time, "CTS 7000 Campo Maior", 0,
		),
		reading.NewReading(
			time, "CTS 6000 Valença do Minho", 1,
		),
	}

	expected := false

	for _, reading := range input {
		actual := base.Equal(reading)
		if expected != actual {
			t.Errorf("expected %t, actual %t", expected, actual)
		}
	}

}

func TestCompareDayEqualDates(t *testing.T) {
	time, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}
	r := reading.NewReading(time, "CTS 6000 Valença do Minho", 0)

	expected := 0
	actual := r.CompareDay(time)

	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestCompareDayHigherDates(t *testing.T) {
	time, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}
	r := reading.NewReading(time, "CTS 6000 Valença do Minho", 0)

	expected := -1
	actual := r.CompareDay(time.AddDate(0, 0, 1))

	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestCompareDayLowerDates(t *testing.T) {
	time, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}
	r := reading.NewReading(time, "CTS 6000 Valença do Minho", 0)

	expected := 1
	actual := r.CompareDay(time.AddDate(0, 0, -1))

	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestCompareReadingDayEqualDates(t *testing.T) {
	time, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}
	r1 := reading.NewReading(time, "CTS 6000 Valença do Minho", 0)
	r2 := reading.NewReading(time, "CTS 6000 Valença do Minho", 0)

	expected := 0
	actual := r1.CompareReadingDay(r2)

	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
