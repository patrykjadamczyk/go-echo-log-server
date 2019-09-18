package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/gen2brain/beeep"
)

func requestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr
		fmt.Println(start, "Received Request", r.Method, r.RequestURI, requesterIP)
		message := fmt.Sprintln(start, "Received Request", r.Method, r.RequestURI, requesterIP)
		beeep.Notify("Go Echo Log Server", message, "")

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
		log.Printf(
		    "DATA: %s",
		    r.Body,
		)
	})
}

func logRoute(w http.ResponseWriter, r *http.Request) {
	html := ""
	w.Write([]byte(html))
}

func main() {
	fileName := "webrequests.log"

	fmt.Println("Making log file ready", "Logfile: ", fileName)
	// https://www.socketloop.com/tutorials/golang-how-to-save-log-messages-to-file
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	defer logFile.Close()

	// direct all log messages to webrequests.log
	log.SetOutput(logFile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", logRoute)

	fmt.Println("Starting Go Logging Server on port 7777")
	http.ListenAndServe(":7777", requestLogger(mux))
}
