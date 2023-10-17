package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	timezone := r.URL.Query().Get("tz")
	timeNow := time.Now()
	if timezone != "" {
		loc, err := time.LoadLocation(timezone)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
		}
		fmt.Fprint(w, timeNow.In(loc))
	} else {
		fmt.Fprint(w, "No timezone supplied!")
	}
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/TimeZoneHttpTrigger", helloHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
