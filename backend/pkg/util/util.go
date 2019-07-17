package util

import "time"

func FormatTime(layout string, dateTime time.Time) (string, error) {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return "", err
	}
	dateTime.In(loc)

	return dateTime.Format(layout), nil

}