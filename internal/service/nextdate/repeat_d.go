package nextdate

import (
	"fmt"
	"strconv"
	"time"

	entityconfig "github.com/probuborka/go_final_project/internal/entity/config"
	entityerror "github.com/probuborka/go_final_project/internal/entity/error"
)

type d struct {
	now  time.Time
	date time.Time
	days int
}

func newD(now time.Time, date time.Time, repeat []string) (date, error) {

	if len(repeat) != 2 {
		return nil, fmt.Errorf("%w: repeat D %s", entityerror.ErrFormatError, repeat)
	}

	days, err := strconv.Atoi(repeat[1])
	if err != nil {
		return nil, fmt.Errorf("%w: not a number %w", entityerror.ErrFormatError, err)
	}

	if days < 1 || days > 400 {
		return nil, fmt.Errorf("%w: %v <> [1..400]", entityerror.ErrNotInInterval, days)
	}

	return d{
		now:  now,
		date: date,
		days: days,
	}, nil
}

func (d d) Next() string {

	nextDate := d.date.AddDate(0, 0, d.days)
	for nextDate.Before(d.now) {
		nextDate = nextDate.AddDate(0, 0, d.days)
	}

	return nextDate.Format(entityconfig.Format1)
}
