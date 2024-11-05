package entity

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const format = "20060102"

var (
	ErrNextData = errors.New("func NextData()")
	ErrRuleD    = errors.New("func ruleD()")
)

func NextDate(nowDate time.Time, date string, repeat string) (string, error) {

	startDate, err := time.Parse(format, date)
	if err != nil {
		return "", fmt.Errorf("%w: date format error", ErrNextData)
	}

	if repeat == "" {
		return "", fmt.Errorf("%w: is empty", ErrNextData)
	}

	repeats := strings.Split(repeat, " ")

	//rules
	switch repeats[0] {
	case "d":
		return ruleD(nowDate, startDate, repeats)
	case "y":
		return ruleY()
	case "w":
		return ruleW()
	case "m":
		return ruleM()
	default:
		return "", fmt.Errorf("%w: rule not found", ErrNextData)
	}
}

// d <число> — задача переносится на указанное число дней. Максимально допустимое число равно 400. Примеры:
// d 1 — каждый день;
// d 7 — для вычисления следующей даты добавляем семь дней;
// d 60 — переносим на 60 дней.
func ruleD(nowDate time.Time, startDate time.Time, repeat []string) (string, error) {

	if len(repeat) != 2 {
		return "", fmt.Errorf("%w: format error", ErrRuleD)
	}

	days, err := strconv.Atoi(repeat[1])
	if err != nil {
		return "", fmt.Errorf("%w: not a number %w", ErrRuleD, err)
	}

	if days < 1 || days > 400 {
		return "", fmt.Errorf("%w: number of days outside the interval (1..400)", ErrRuleD)
	}

	nextDate := startDate.AddDate(0, 0, days)
	for nextDate.Before(nowDate) {
		nextDate = nextDate.AddDate(0, 0, days)
	}

	return nextDate.Format(format), nil
}

func ruleY() (string, error) {
	return "", nil
}

func ruleW() (string, error) {
	return "", nil
}

func ruleM() (string, error) {
	return "", nil
}
