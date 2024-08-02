package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/application/usecase"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/application/usecase/pcs"
)

const timelayout = "2006-01-02"

type PcsGetController struct {
	usecase usecase.UseCase[pcs.AverageProps, float64]
}

func NewPcsGetController(
	usecase usecase.UseCase[pcs.AverageProps, float64],
) PcsGetController {
	return PcsGetController{
		usecase,
	}
}

func (c PcsGetController) Handle(r *http.Request) (string, error) {
	startTimeStr := r.URL.Query().Get("startTime")
	if len(startTimeStr) == 0 {
		startTimeStr = time.Now().Format(timelayout)
	}

	startTime, err := time.Parse(timelayout, startTimeStr)
	if err != nil {
		return "", err
	}
	stopTimeStr := r.URL.Query().Get("stopTime")
	stopTime := startTime
	if len(stopTimeStr) > 0 {
		if stopTime, err = time.Parse(timelayout, stopTimeStr); err != nil {
			return "", err
		}
	}
	input := pcs.NewAverageProps(
		r.Context(),
		startTime.Add(time.Hour*5),
		stopTime.Add(time.Hour*4),
	)
	average, err := c.usecase.Execute(input)
	if err != nil {
		return "", err
	}
	return strconv.FormatFloat(average, 'g', -1, 64), nil
}
