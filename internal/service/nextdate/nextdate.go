package nextdate

import (
	"fmt"
	"strings"
	"time"

	"github.com/probuborka/go_final_project/internal/entity"
)

type date interface {
	Next() string
}

func New(now time.Time, dateStr string, repeat string) (date, error) {
	date, err := time.Parse(entity.Format1, dateStr)
	if err != nil {
		return nil, err
	}

	repeats := strings.Split(repeat, " ")

	//rules
	switch repeats[0] {
	case "d":
		return newD(now, date, repeats)
	case "y":
		return newY(now, date, repeats)
	case "w":
		return newW(now, date, repeats)
	case "m":
		return newM(now, date, repeats)
	default:
		return nil, fmt.Errorf("%w: repeat %s", entity.ErrNotFound, repeats[0])
	}
}
