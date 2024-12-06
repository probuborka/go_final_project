package nextdate

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	entityconfig "github.com/probuborka/go_final_project/internal/entity/config"
	entityerror "github.com/probuborka/go_final_project/internal/entity/error"
)

type w struct {
	now  time.Time
	date time.Time
	days map[int]struct{}
}

func newW(now time.Time, date time.Time, repeat []string) (date, error) {

	if len(repeat) != 2 {
		return nil, fmt.Errorf("%w: repeat W %s", entityerror.ErrFormatError, repeat)
	}

	daysWeeks := strings.Split(repeat[1], ",")

	if len(daysWeeks) > 7 {
		return nil, fmt.Errorf("%w: days in week %v", entityerror.ErrFormatError, len(daysWeeks))
	}

	days := make(map[int]struct{})
	for _, v := range daysWeeks {
		day, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("%w: not a number %w", entityerror.ErrFormatError, err)
		}

		if day < 1 || day > 7 {
			return nil, fmt.Errorf("%w: %v <> [1..7]", entityerror.ErrNotInInterval, day)
		}

		//
		days[day] = struct{}{}
	}

	return w{
		now:  now,
		date: date,
		days: days,
	}, nil
}

func (w w) Next() string {

	if w.date.Before(w.now) {
		w.date = w.now
	}

	nextDate := w.date
	for d := 1; d <= 7; d++ {
		nextDate = nextDate.AddDate(0, 0, 1)
		day := int(nextDate.Weekday())
		if day == 0 {
			day = 7
		}
		if _, ok := w.days[day]; ok {
			break
		}
	}
	return nextDate.Format(entityconfig.Format1)
}
