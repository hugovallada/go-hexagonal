package converter

import (
	"time"
)

const (
	FORMAT = "02/01/2006"
)

func StringToDateConverter(dateAsString string) (time.Time, error) {
	date, err := time.Parse(FORMAT, dateAsString)
	if err != nil {
		return time.Now(), err
	}
	return date, nil
}
