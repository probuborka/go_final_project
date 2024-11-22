package tests

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNextDateMy(t *testing.T) {

	tbl := []nextDate{
		{"20240409", "m 31", "20240531"},
		{"20240329", "m 10,17 12,8,1", "20240810"},
		{"20231106", "m 13", "20240213"},
		{"20240120", "m 40,11,19", ""},
		{"20240116", "m 16,5", "20240205"},
		{"20240126", "m 25,26,7", "20240207"},
		{"20230311", "m 07,19 05,6", "20240507"},
		{"20230311", "m 1 1,2", "20240201"},
		{"20240127", "m -1", "20240131"},
		{"20240222", "m -2", "20240228"},
		{"20240222", "m -2,-3", ""},
		{"20240326", "m -1,-2", "20240330"},
		{"20240201", "m -1,18", "20240218"},
	}

	check := func() {
		for _, v := range tbl {
			urlPath := fmt.Sprintf("api/nextdate?now=20240126&date=%s&repeat=%s",
				url.QueryEscape(v.date), url.QueryEscape(v.repeat))
			get, err := getBody(urlPath)
			assert.NoError(t, err)
			next := strings.TrimSpace(string(get))
			_, err = time.Parse("20060102", next)
			if err != nil && len(v.want) == 0 {
				continue
			}
			assert.Equal(t, v.want, next, `{%q, %q, %q}`,
				v.date, v.repeat, v.want)
		}
	}

	check()
}
