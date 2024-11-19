package nextdate

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/probuborka/go_final_project/internal/entity"
)

type w struct {
	now  time.Time
	date time.Time
	days map[int]struct{}
}

func newW(now time.Time, date time.Time, repeat []string) (date, error) {

	if len(repeat) != 2 {
		return nil, fmt.Errorf("%w: rule W %s", entity.ErrFormatError, repeat)
	}

	daysWeeks := strings.Split(repeat[1], ",")

	if len(daysWeeks) > 7 {
		return nil, fmt.Errorf("%w: days in week %v", entity.ErrFormatError, len(daysWeeks))
	}

	days := make(map[int]struct{})
	for _, v := range daysWeeks {
		day, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("%w: not a number %w", entity.ErrFormatError, err)
		}

		if day < 1 || day > 7 {
			return nil, fmt.Errorf("%w: %v <> [1..7]", entity.ErrNotInInterval, day)
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
	return nextDate.Format(entity.Format)
}
