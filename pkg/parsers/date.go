package parsers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ConvertDate(date string) (time.Time, time.Time, error) {
	result := strings.Split(date, "/")
	if len(result) != 3 {
		return time.Time{}, time.Time{},
			fmt.Errorf("Unable to parse date: %s", date)
	}

	y, err := strconv.ParseInt(result[2], 10, 64)
	if err != nil {
		return time.Time{}, time.Time{},
			fmt.Errorf("Unable to parse date: %s", date)
	}

	m, err := strconv.ParseInt(result[1], 10, 64)
	if err != nil {
		return time.Time{}, time.Time{},
			fmt.Errorf("Unable to parse date: %s", date)
	}

	d, err := strconv.ParseInt(result[0], 10, 64)
	if err != nil {
		return time.Time{}, time.Time{},
			fmt.Errorf("Unable to parse date: %s", date)
	}

	start := time.Date(int(y), time.Month(m), int(d), 0, 0, 0, 0, time.UTC)
	return start, start.Add(time.Hour * 24), nil
}
