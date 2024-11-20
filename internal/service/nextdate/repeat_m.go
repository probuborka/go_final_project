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

	switch {
	case m.len == 2:
		if m.date.Before(m.now) {
			m.date = m.now
		}
		for {
			//
			nextDate := m.date
			// current day
			curDay := m.date.Day()
			// last day
			lastDay := m.date.AddDate(0, 1, -curDay).Day()
			//
			day := 0
			m1 := 0
			for _, v := range m.days {
				if v < 0 {
					day = lastDay + v + 1
				} else {
					day = v
				}

				if day > curDay && day <= lastDay-m1 {
					nextDate = m.date.AddDate(0, 0, day-curDay)
					if v == -2 {
						m1 = 1
					}
					if v > 0 {
						break
					}
				}
			}
			if nextDate.After(m.date) {
				return nextDate.Format(entity.Format)
			}
			m.date = m.date.AddDate(0, 1, -curDay+1)
		}
	case m.len == 3:
		if m.date.Before(m.now) {
			m.date = m.now
		}
		startDateCheck := m.date
		for {
			//
			nextDate := m.date
			// current day
			curDay := m.date.Day()
			// last day
			lastDay := m.date.AddDate(0, 1, -curDay).Day()
			// current month
			curMonth := int(m.date.Month())
			for _, m1 := range m.months {
				if m1 < curMonth {
					continue
				} else if m1 > curMonth {
					m.date = m.date.AddDate(0, m1-curMonth, -curDay+1)
					//
					startDateCheck = m.date.AddDate(0, 0, -1)
					//
					nextDate = m.date
					// current day
					curDay = m.date.Day()
					// last day
					lastDay = m.date.AddDate(0, 1, -curDay).Day()
					//
					curMonth = int(m.date.Month())
				}
				//
				day := 0
				m1 := 0
				for _, d := range m.days {
					if d < 0 {
						day = lastDay + d + 1
					} else {
						day = d
					}

					if day >= curDay && day <= lastDay-m1 {
						nextDate = m.date.AddDate(0, 0, day-curDay)
						if d == -2 {
							m1 = 1
						}
						if d > 0 {
							break
						}
					}
				}
				if nextDate.After(startDateCheck) {
					return nextDate.Format(entity.Format)
				}
			}
			m.date = m.date.AddDate(0, 12-curMonth+1, -curDay+1)
			startDateCheck = m.date.AddDate(0, 0, -1)
		}
	default:
		return "" //, fmt.Errorf("%w: :-(")
	}
}
