package logserver

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
	"github.com/nu7hatch/gouuid"
	"github.com/patrykjadamczyk/go-echo-log-server/config"
)

func banner(appConfiguration config.Configuration) {
	fmt.Println("Go Echo Log Server")
	fmt.Println("Version:", appConfiguration.App.Version)
}

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
		errNotify := beeep.Notify("Go Echo Log Server", message, "")
		if errNotify != nil {
			fmt.Println(errNotify)
		}
		log.Printf(message)
	})
}

func logRoute(w http.ResponseWriter, r *http.Request) {
	requestIdentifier, errUUID := uuid.NewV4()
	if errUUID != nil {
		fmt.Println(errUUID)
		return
	}
	html := fmt.Sprintf("%s", requestIdentifier.String())
	w.WriteHeader(http.StatusAccepted)
	_, err := w.Write([]byte(html))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Main(appConfiguration config.Configuration) {
	banner(appConfiguration)
	fileName := appConfiguration.LogServer.LogFile
	port := appConfiguration.LogServer.Port

	fmt.Println("Preparing Log File...")
	fmt.Println("Logfile prepared. Logfile location: ", fileName)
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	defer logFile.Close()

	// direct all log messages to log file
	log.SetOutput(logFile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", logRoute)

	fmt.Println("Starting Go Logging Server on port", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), requestLogger(mux))
}
