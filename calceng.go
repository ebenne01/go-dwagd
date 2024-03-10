// Copyright 2024 Edward Bennett.  All rights reserved.
// Use of this source code is governed by an Apache 2.0
// style license that can be found in the LICENSE file.

package main

import (
	"errors"
	"strconv"
	"time"
)

type month struct {
	Name      string
	KeyNum    int
	LeapYrAdj int
}

var months = map[int]month{
	1:  {"January", 1, -1},
	2:  {"February", 4, -1},
	3:  {"March", 4, 0},
	4:  {"April", 0, 0},
	5:  {"May", 2, 0},
	6:  {"June", 5, 0},
	7:  {"July", 0, 0},
	8:  {"August", 3, 0},
	9:  {"September", 6, 0},
	10: {"October", 1, 0},
	11: {"November", 4, 0},
	12: {"December", 6, 0},
}

var days = map[int]string{
	0: "Saturday",
	1: "Sunday",
	2: "Monday",
	3: "Tuesday",
	4: "Wednesday",
	5: "Thursday",
	6: "Friday",
}

func calculateDayOfWeek(tm time.Time) (string, error) {

	monthVal := int(tm.Month())
	dayVal := tm.Day()
	yearVal := tm.Year()
	shortYear, err := strconv.Atoi(tm.Format(twoDigitYear))

	if yearVal < 1753 {
		return "", errors.New("year must be greater than 1752")
	}

	if err != nil {
		return "", errors.New("bad year")
	}

	var sum = shortYear
	sum += shortYear / 4
	sum += dayVal

	monthKey := months[monthVal]
	sum += monthKey.KeyNum

	if isLeapYear(yearVal) {
		sum += monthKey.LeapYrAdj
	}

	if yearVal < 1800 {
		sum += 4
	} else if yearVal < 1900 {
		sum += 2
	} else if yearVal > 1999 {
		sum -= 1
	}

	dayNum := sum % 7
	return days[dayNum], nil
}

func isLeapYear(year int) bool {
	return (year%400 == 0) || ((year%4 == 0) && year%100 != 0)
}
