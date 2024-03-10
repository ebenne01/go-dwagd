// Copyright 2024 Edward Bennett.  All rights reserved.
// Use of this source code is governed by an Apache 2.0
// style license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func badRequestHandler(w http.ResponseWriter, r *http.Request) {
	msg := map[string]any{"error": http.StatusText(http.StatusBadRequest),
		"message": "Invalid path",
		"path":    r.URL.Path}
	writeJSONResponse(w, http.StatusBadRequest, msg)
}

func dayOfWeekHandler(w http.ResponseWriter, r *http.Request) {
	dateParam := r.PathValue("date")
	tm, err := time.Parse(layoutISO, dateParam)

	if err != nil {
		msg := map[string]any{"error": http.StatusText(http.StatusBadRequest),
			"message": "Invalid date format.  Must be yyyy-mm-dd"}
		writeJSONResponse(w, http.StatusBadRequest, msg)
		return
	}

	day, err := calculateDayOfWeek(tm)

	if err != nil {
		msg := map[string]any{"error": http.StatusText(http.StatusBadRequest),
			"message": err.Error()}
		writeJSONResponse(w, http.StatusBadRequest, msg)
		return
	}

	writeJSONResponse(w, http.StatusOK, day)
}

// helper functions
func writeJSONResponse(w http.ResponseWriter, code int, obj any) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	jsonBytes, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}
