// Copyright 2022 Edward Bennett.  All rights reserved.
// Use of this source code is governed by an Apache 2.0
// style license that can be found in the LICENSE file.

package main

import (
	"net/http"
)

const (
	layoutISO    = "2006-01-02"
	twoDigitYear = "06"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/dayofweek/{date}", dayOfWeekHandler)
	mux.HandleFunc("/", badRequestHandler)
	http.ListenAndServe(":8080", mux)
}
