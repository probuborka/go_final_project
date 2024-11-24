package nextdate

import (
	"fmt"
	"time"

	entityconfig "github.com/probuborka/go_final_project/internal/entity/config"
	entityerror "github.com/probuborka/go_final_project/internal/entity/error"
)

type y struct {
	now  time.Time
	date time.Time
}

func newY(now time.Time, date time.Time, repeat []string) (date, error) {
	if len(repeat) != 1 {
		return nil, fmt.Errorf("%w: repeat Y %s", entityerror.ErrFormatError, repeat)
	}

	return y{
		now:  now,
		date: date,
	}, nil
}

func (y y) Next() string {

	nextDate := y.date.AddDate(1, 0, 0)
	for nextDate.Before(y.now) {
		nextDate = nextDate.AddDate(1, 0, 0)
	}

	return nextDate.Format(entityconfig.Format1)
}
