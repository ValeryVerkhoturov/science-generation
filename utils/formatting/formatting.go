package formatting

import (
	"time"
)

func FormatDate(date string) (string, error) {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return "", err
	}
	return t.Format("02.01.2006"), nil
}
