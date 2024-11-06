package entity

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	errNextData = errors.New("func NextData()")
	errRuleD    = errors.New("func ruleD()")
	errRuleY    = errors.New("func ruleY()")
	errRuleW    = errors.New("func ruleW()")
	errRuleM    = errors.New("func ruleM()")
)

func NextDate(nowDate time.Time, date string, repeat string) (string, error) {
	//check
	startDate, err := time.Parse(Format, date)
	if err != nil {
		return "", fmt.Errorf("%w: date format error", errNextData)
	}

	if repeat == "" {
		return "", fmt.Errorf("%w: is empty", errNextData)
	}

	repeats := strings.Split(repeat, " ")

	//rules
	switch repeats[0] {
	case "d":
		return ruleD(nowDate, startDate, repeats)
	case "y":
		return ruleY(nowDate, startDate, repeats)
	case "w":
		return ruleW(nowDate, repeats)
	case "m":
		return ruleM(nowDate, startDate, repeats)
	default:
		return "", fmt.Errorf("%w: rule not found", errNextData)
	}
}

// d <число> — задача переносится на указанное число дней. Максимально допустимое число равно 400. Примеры:
// d 1 — каждый день;
// d 7 — для вычисления следующей даты добавляем семь дней;
// d 60 — переносим на 60 дней.
func ruleD(nowDate time.Time, startDate time.Time, repeat []string) (string, error) {
	//check
	if len(repeat) != 2 {
		return "", fmt.Errorf("%w: format error", errRuleD)
	}

	days, err := strconv.Atoi(repeat[1])
	if err != nil {
		return "", fmt.Errorf("%w: not a number %w", errRuleD, err)
	}

	if days < 1 || days > 400 {
		return "", fmt.Errorf("%w: number of days outside the interval (1..400)", errRuleD)
	}

	//calculations
	nextDate := startDate.AddDate(0, 0, days)
	for nextDate.Before(nowDate) {
		nextDate = nextDate.AddDate(0, 0, days)
	}

	return nextDate.Format(Format), nil
}

// y — задача выполняется ежегодно. Этот параметр не требует дополнительных уточнений.
// При выполнении задачи дата перенесётся на год вперёд.
func ruleY(nowDate time.Time, startDate time.Time, repeat []string) (string, error) {
	//check
	if len(repeat) != 1 {
		return "", fmt.Errorf("%w: format error", errRuleY)
	}

	//calculations
	nextDate := startDate.AddDate(1, 0, 0)
	for nextDate.Before(nowDate) {
		nextDate = nextDate.AddDate(1, 0, 0)
	}

	return nextDate.Format(Format), nil
}

// w <через запятую от 1 до 7> — задача назначается в указанные дни недели, где 1 — понедельник, 7 — воскресенье. Например:
// w 7 — задача перенесётся на ближайшее воскресенье;
// w 1,4,5 — задача перенесётся на ближайший понедельник, четверг или пятницу;
// w 2,3 — задача перенесётся на ближайший вторник или среду.
func ruleW(nowDate time.Time, repeat []string) (string, error) {
	//check
	if len(repeat) != 2 {
		return "", fmt.Errorf("%w: format error", errRuleW)
	}

	daysWeeks := strings.Split(repeat[1], ",")

	if len(daysWeeks) > 7 {
		return "", fmt.Errorf("%w: format error", errRuleW)
	}

	deys := make(map[int]struct{})
	for _, v := range daysWeeks {
		day, err := strconv.Atoi(v)
		if err != nil {
			return "", fmt.Errorf("%w: not a number %w", errRuleW, err)
		}

		if day < 1 || day > 7 {
			return "", fmt.Errorf("%w: number of days outside the interval (1..7)", errRuleW)
		}

		//
		deys[day] = struct{}{}
	}

	//calculations
	nextDate := nowDate
	for d := 1; d <= 7; d++ {
		nextDate = nextDate.AddDate(0, 0, 1)
		day := int(nextDate.Weekday())
		if day == 0 {
			day = 7
		}
		if _, ok := deys[day]; ok {
			break
		}
	}
	return nextDate.Format(Format), nil
}

// m <через запятую от 1 до 31,-1,-2> [через запятую от 1 до 12] — задача назначается в указанные дни месяца.
// При этом вторая последовательность чисел опциональна и указывает на определённые месяцы. Например:
// m 4 — задача назначается на четвёртое число каждого месяца;
// m 1,15,25 — задача назначается на 1-е, 15-е и 25-е число каждого месяца;
// m -1 — задача назначается на последний день месяца;
// m -2 — задача назначается на предпоследний день месяца;
// m 3 1,3,6 — задача назначается на 3-е число января, марта и июня;
// m 1,-1 2,8 — задача назначается на 1-е и последнее число февраля и авгуcта.
func ruleM(nowDate time.Time, startDate time.Time, repeat []string) (string, error) {
	//check
	len := len(repeat)
	if len != 2 && len != 3 {
		return "", fmt.Errorf("%w: format error", errRuleM)
	}

	//check day
	days := strings.Split(repeat[1], ",")
	for _, v := range days {
		day, err := strconv.Atoi(v)
		if err != nil {
			return "", fmt.Errorf("%w: not a number %w", errRuleM, err)
		}

		if day < -2 || day > 31 {
			return "", fmt.Errorf("%w: number of day outside the interval (-2..31)", errRuleM)
		}
	}

	sort.Slice(days, func(i, j int) bool {
		return days[i] < days[j]
	})

	//check month
	months := make([]string, 0)
	if len == 3 {
		months = strings.Split(repeat[2], ",")
		for _, v := range months {
			month, err := strconv.Atoi(v)
			if err != nil {
				return "", fmt.Errorf("%w: not a number %w", errRuleM, err)
			}

			if month < 1 || month > 12 {
				return "", fmt.Errorf("%w: number of month outside the interval (1..12)", errRuleM)
			}
		}

		sort.Slice(months, func(i, j int) bool {
			return months[i] < months[j]
		})
	}

	//calculations
	switch len {
	case 2:
		return "", nil
	case 3:
		return "", nil
	default:
		return "", fmt.Errorf("%w: :-(", errRuleM)
	}
}
