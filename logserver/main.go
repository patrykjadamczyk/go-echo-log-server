package logserver

import (
	"encoding/json"
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
	"github.com/patrykjadamczyk/go-echo-log-server/logdb"
)

func banner(appConfiguration config.Configuration) {
	fmt.Println("Go Echo Log Server")
	fmt.Println("Version:", appConfiguration.App.Version)
}

func requestLogger(targetMux http.Handler, appDatabase logdb.DB) http.Handler {
	// Handle Requests
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Time when Request Arrived
		start := time.Now()
		// Make Request Identifier
		requestIdentifier, errUUID := uuid.NewV4()
		if errUUID != nil {
			fmt.Println(errUUID)
		}
		// Serve Route on HTTP
		targetMux.ServeHTTP(w, r)
		// Make pretty dump of Request
		requestDump, err := httputil.DumpRequest(r, false)
		if err != nil {
			fmt.Println(err)
		}
		// Get Body from buffer
		bodyBuffer, _ := ioutil.ReadAll(r.Body)
		body, _ := url.QueryUnescape(string(bodyBuffer))
		// Parse Form Data
		_ = r.ParseForm()
		// Make Message
		message := fmt.Sprintln(start) + fmt.Sprintln(string(requestDump)) + fmt.Sprintln(body)
		// Log Message to STDOUT of Server
		fmt.Println(message)
		// Log Message to system notifications
		errNotify := beeep.Notify(
			fmt.Sprintf("GELS: %s", requestIdentifier.String()),
			message,
			"")
		if errNotify != nil {
			fmt.Println(errNotify)
		}
		// Log Message to Log File
		log.Printf(message)
		// Add request data to memory database
		_ = appDatabase.Update(func(tx *logdb.Tx) error {
			jsonData, _ := json.Marshal(r.Form)
			tx.Set(requestIdentifier.String(), logdb.LogRequest{
				Start:           fmt.Sprint(start),
				RequestInfo:     string(requestDump),
				RequestDataForm: string(jsonData),
				RequestData:     fmt.Sprint(body),
			})
			return nil
		})
	})
}

func logRoute(w http.ResponseWriter, r *http.Request) {
	html := ""
	w.WriteHeader(http.StatusAccepted)
	_, err := w.Write([]byte(html))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Main(appConfiguration config.Configuration, appDatabase logdb.DB) {
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
	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), requestLogger(mux, appDatabase))
}
