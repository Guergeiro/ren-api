package entity

import (
	"cmp"
	"time"
)

type Reading struct {
	timestamp time.Time
	name      MonitorizationPoint
	value     float64
}

type MonitorizationPoint string

const (
	VALENCA_MINHO MonitorizationPoint = "CTS 6000 Valença do Minho"
	CAMPO_MAIOR   MonitorizationPoint = "CTS 7000 Campo Maior"
)

func NewReading(
	timestamp time.Time,
	name string,
	value float64,
) Reading {

	return Reading{
		timestamp: timestamp,
		name:      MonitorizationPoint(name),
		value:     value,
	}
}

func (r Reading) Timestamp() time.Time {
	return r.timestamp
}

func (r Reading) Name() MonitorizationPoint {
	return r.name
}

func (r Reading) Value() float64 {
	return r.value
}

func (r1 Reading) Equal(r2 Reading) bool {
	if r1.timestamp.Equal(r2.timestamp) == false {
		return false
	}

	if r1.name != r2.name {
		return false
	}

	return cmp.Compare(r1.value, r2.value) == 0
}

func (r1 Reading) CompareReadingDay(r2 Reading) int {
	return r1.CompareDay(r2.timestamp)
}

func (r Reading) CompareDay(date time.Time) int {
	baseReadingDate := time.Date(
		r.timestamp.Year(),
		r.timestamp.Month(),
		r.timestamp.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)

	baseDate := time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)

	return baseReadingDate.Compare(baseDate)
}
