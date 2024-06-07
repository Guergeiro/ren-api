package interval

import (
	"fmt"
	"testing"
	"time"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/interval"
)

const timelayout = "2006-01-02"

func TestEqualInterval(t *testing.T) {
	start, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}
	stop := start.AddDate(0, 0, 10)

	i1 := interval.NewInterval(start, stop)
	i2 := interval.NewInterval(start, stop)

	expected := true
	actual := i1.Equal(i2)

	if expected != actual {
		t.Errorf("expected %t, actual %t", expected, actual)
	}
}

func TestDifferentInterval(t *testing.T) {
	start, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}
	stop := start.AddDate(0, 0, 10)

	base := interval.NewInterval(start, stop)

	input := []interval.Interval{
		interval.NewInterval(start.AddDate(0, 0, 1), stop),
		interval.NewInterval(start.AddDate(0, 0, -1), stop),
		interval.NewInterval(start, stop.AddDate(0, 0, 1)),
		interval.NewInterval(start, stop.AddDate(0, 0, -1)),
	}

	expected := false
	for _, interval := range input {
		actual := base.Equal(interval)
		if expected != actual {
			t.Errorf("expected %t, actual %t", expected, actual)
		}
	}
}


func TestNewIntervalsNoExcess(t *testing.T) {
	start, err := time.Parse(timelayout, "2024-05-22")
	if err != nil {
		t.Fail()
	}
	stop := start.AddDate(0, 1, 0)

	output := interval.NewIntervals(start, stop)

	expected := 1
	actual := len(output)

	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestNewIntervalsSingleExcess(t *testing.T) {
	start, err := time.Parse(timelayout, "2024-04-01")
	if err != nil {
		t.Fail()
	}
	stop := start.AddDate(0, 1, 1)

	output := interval.NewIntervals(start, stop)

	expected := 2
	actual := len(output)
	fmt.Println(output)

	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestNewIntervalsDoubleExcess(t *testing.T) {
	start, err := time.Parse(timelayout, "2024-04-01")
	if err != nil {
		t.Fail()
	}
	stop := start.AddDate(0, 2, 2)

	output := interval.NewIntervals(start, stop)

	expected := 3
	actual := len(output)
	fmt.Println(output)

	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
