package tsql

import (
	. "github.com/tbshill/goparsec"
	"strconv"
	"time"
)

var _ Node = &DateStringNode{}

// DateStringNode represents a date, inside the '' of a sql statement
type DateStringNode struct {
	Date time.Time
}

func (d *DateStringNode) Parse(in string) (string, error) {
	ExpectYear := And(ExpectDigit, ExpectDigit, ExpectDigit, ExpectDigit)
	ExpectMonth := And(ExpectDigit, ExpectDigit)
	ExpectDay := And(ExpectDigit, ExpectDigit)

	var yearStr, monthStr, dayStr, rem string
	var year, month, day int
	var err error

	if yearStr, rem, err = ExpectYear(in); err != nil {
		return in, err
	}

	if _, rem, err = ExpectMinus(rem); err != nil {
		return in, err
	}

	if monthStr, rem, err = ExpectMonth(rem); err != nil {
		return in, err
	}

	if dayStr, rem, err = ExpectDay(rem); err != nil {
		return in, err
	}

	if year, err = strconv.Atoi(yearStr); err != nil {
		return in, err
	}

	if month, err = strconv.Atoi(monthStr); err != nil {
		return in, err
	}

	if day, err = strconv.Atoi(dayStr); err != nil {
		return in, err
	}

	d.Date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	return rem, nil
}

func (d *DateStringNode) String() string {
	year, month, day := d.Date.Date()
	yearStr := strconv.Itoa(year)
	monthStr := strconv.Itoa(int(month))
	dayStr := strconv.Itoa(day)
	return yearStr + "-" + monthStr + "-" + dayStr
}
