package nextdate

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/probuborka/go_final_project/internal/entity"
)

type m struct {
	now    time.Time
	date   time.Time
	len    int
	days   []int
	months []int
}

func newM(now time.Time, date time.Time, repeat []string) (date, error) {
	len := len(repeat)
	if len != 2 && len != 3 {
		return nil, fmt.Errorf("%w: repeat M %s", entity.ErrFormatError, repeat)
	}

	//check day
	strDays := strings.Split(repeat[1], ",")
	days := make([]int, 0)
	for _, v := range strDays {
		day, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("%w: not a number %w", entity.ErrFormatError, err)
		}

		if day < -2 || day > 31 {
			return nil, fmt.Errorf("%w: %v <> [-2..31]", entity.ErrNotInInterval, day)
		}

		//
		days = append(days, day)
	}

	sort.Slice(days, func(i, j int) bool {
		return days[i] < days[j]
	})

	//check month
	months := make([]int, 0)
	if len == 3 {
		strMonths := strings.Split(repeat[2], ",")
		for _, v := range strMonths {
			month, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("%w: not a number %w", entity.ErrFormatError, err)
			}

			if month < 1 || month > 12 {
				return nil, fmt.Errorf("%w: %v <> [1..12]", entity.ErrNotInInterval, month)
			}

			//
			months = append(months, month)
		}

		sort.Slice(months, func(i, j int) bool {
			return months[i] < months[j]
		})
	}

	return m{
		now:    now,
		date:   date,
		len:    len,
		days:   days,
		months: months,
	}, nil
}

func (m m) Next() string {
	var months []int
	if m.len == 3 {
		months = append(months, m.months...)
	} else {
		months = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	}
	date := m.date
	if date.Before(m.now) {
		date = m.now
	}

	ys, mm, ds := date.Date()
	ms := int(mm)
	yss := ys
	for i := 0; i < 2; i++ {
		for _, v := range months {
			if v < ms {
				continue
			}
			if v != ms {
				ds = 1
			} else if v == ms && ys == yss {
				ds++
			}
			ms = v
			days := daysInMonths(m.days, ys, ms, date.Location())
			for _, v := range days {
				if v < ds {
					continue
				}
				ds = v
				date := time.Date(ys, time.Month(ms), ds, 0, 0, 0, 0, date.Location())
				if time.Month(ms) == date.Month() {
					return date.Format(entity.Format1)
				}
			}
		}
		ys++
	}
	return ""
}

func daysInMonths(days []int, ys, ms int, location *time.Location) []int {
	retDays := make([]int, 0)
	date := time.Date(ys, time.Month(ms), 1, 0, 0, 0, 0, location)
	lastOfMonth := date.AddDate(0, 1, -1)
	for _, v := range days {
		if v > 0 {
			retDays = append(retDays, v)
		} else if v == -1 {
			retDays = append(retDays, lastOfMonth.Day())
		} else if v == -2 {
			retDays = append(retDays, lastOfMonth.Day()-1)
		}
	}

	sort.Slice(retDays, func(i, j int) bool {
		return retDays[i] < retDays[j]
	})

	return retDays
}
