package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func writeResponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = message

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error in JSON marshal: %s", err)
	}

	w.Write(jsonResp)
}

func timezoneHandler(w http.ResponseWriter, r *http.Request) {
	timezone := r.URL.Query().Get("tz")
	timeNow := time.Now()
	if timezone == "" {
		writeResponse(w, "No timezone supplied!", http.StatusBadRequest)
		return
	}
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatalf("Error in loading timezone location: %v", err)
	}

	writeResponse(w, timeNow.In(loc).String(), http.StatusOK)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/TimeZoneHttpTrigger", timezoneHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
