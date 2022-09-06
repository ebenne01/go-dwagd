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
	"testing"
	"time"
)

var leapYears = [...]int{
	1804, 1808, 1812, 1816, 1820, 1824, 1832, 1836, 1840, 1844,
	1848, 1852, 1856, 1860, 1864, 1868, 1872, 1876, 1880, 1884,
	1888, 1892, 1896, 1904, 1908, 1912, 1916, 1920, 1924, 1928,
	1932, 1936, 1940, 1944, 1948, 1952, 1956, 1960, 1964, 1968,
	1972, 1976, 1980, 1984, 1988, 1992, 1996, 2000, 2004, 2008,
	2012, 2016, 2020, 2024, 2028, 2032, 2036, 2040, 2044, 2048,
	2052, 2056, 2060, 2064, 2068, 2072, 2076, 2080, 2084, 2088,
	2092, 2096,
}

var nonLeapYears = [...]int{
	1805, 1809, 1813, 1817, 1821, 1825, 1833, 1837, 1841,
	1845, 1849, 1853, 1857, 1861, 1865, 1869, 1873, 1877,
	1881, 1885, 1889, 1893, 1897, 1905, 1909, 1913, 1917,
	1921, 1925, 1929, 1933, 1937, 1941, 1945, 1949, 1953,
	1957, 1961, 1965, 1969, 1973, 1977, 1981, 1985, 1989,
	1993, 1997, 2001, 2005, 2009, 2013, 2017, 2021, 2025,
	2029, 2033, 2037, 2041, 2045, 2049, 2053, 2057, 2061,
	2065, 2069, 2073, 2077, 2081, 2085, 2089, 2093, 2097,
}

var daysOfWeek = map[string]string{
	"1804-12-23": "Sunday",
	"1753-01-01": "Monday",
	"2019-10-29": "Tuesday",
	"1937-08-18": "Wednesday",
	"2037-04-02": "Thursday",
	"1949-10-21": "Friday",
	"1992-06-20": "Saturday",
}

func TestIsLeapYearTrue(t *testing.T) {
	for _, n := range leapYears {
		if isLeapYear(n) == false {
			t.Errorf("%d is not a leap year", n)
		}
		isLeapYear(n)
	}
}

func TestIsLeapYearFalse(t *testing.T) {
	for _, n := range nonLeapYears {
		if isLeapYear(n) == true {
			t.Errorf("%d is a leap year", n)
		}
	}
}

func TestValidDates(t *testing.T) {
	for k, v := range daysOfWeek {
		tm := parseDate(k)
		actual, err := calculateDayOfWeek(tm)
		if err != nil || actual != v {
			t.Errorf("Incorrect day calculated for date %s", k)
		}
	}
}

//
// helper functions
//
func parseDate(s string) time.Time {
	tm, _ := time.Parse(layoutISO, s)
	return tm
}
