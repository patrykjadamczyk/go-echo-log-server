package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

func requestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)
		requestDump, err := httputil.DumpRequest(r, false)
		if err != nil {
			fmt.Println(err)
		}
		bodyBuffer, _ := ioutil.ReadAll(r.Body)
		body, _ := url.QueryUnescape(string(bodyBuffer))

		message := fmt.Sprintln(start) + fmt.Sprintln(string(requestDump)) + fmt.Sprintln(body)
		fmt.Println(message)
		_ = beeep.Notify("Go Echo Log Server", message, "")
		log.Printf(message)
	})
}

func logRoute(w http.ResponseWriter, r *http.Request) {
	html := ""
	_, _ = w.Write([]byte(html))
}

func main() {
	fileName := "webrequests.log"

	fmt.Println("Making log file ready", "Logfile: ", fileName)
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
	_ = http.ListenAndServe(":7777", requestLogger(mux))
}
