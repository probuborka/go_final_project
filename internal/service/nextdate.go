package service

import (
	"time"

	"github.com/probuborka/go_final_project/internal/service/nextdate"
)

func (t task) NextDate(nowDate time.Time, dateStr string, repeat string) (string, error) {

	date, err := nextdate.New(nowDate, dateStr, repeat)
	if err != nil {
		return "", err
	}

	return date.Next(), nil
}
