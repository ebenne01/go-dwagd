// Copyright 2022 Edward Bennett
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	layoutISO    = "2006-01-02"
	twoDigitYear = "06"
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

func main() {
	router := gin.Default()
	router.GET("/dayofweek/:date", getDayOfWeek)
	router.Run("localhost:8080")
}

func getDayOfWeek(c *gin.Context) {
	tm, err := time.Parse(layoutISO, c.Param("date"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid date format.  Must be yyyy-mm-dd")
		return
	}

	day, err := calculateDayOfWeek(tm)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, day)

}

func calculateDayOfWeek(tm time.Time) (string, error) {

	monthVal := int(tm.Month())
	dayVal := tm.Day()
	yearVal := tm.Year()
	shortYear, err := strconv.Atoi(tm.Format(twoDigitYear))

	if yearVal < 1753 {
		return "", errors.New("Year must be greater than 1752")
	}

	if err != nil {
		return "", errors.New("Bad year")
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
