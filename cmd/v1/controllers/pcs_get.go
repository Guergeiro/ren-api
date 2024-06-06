package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/usecases/pcs"
)

const timelayout = "2006-01-02"

func PcsGet(r *http.Request) (string, error) {
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
	average, err := pcs.Average(startTime, stopTime)
	if err != nil {
		return "", err
	}
	return strconv.FormatFloat(average, 'g', -1, 64), nil
}
