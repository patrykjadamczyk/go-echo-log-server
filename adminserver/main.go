package adminserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/patrykjadamczyk/go-echo-log-server/adminserver/adminpanel/statik"
	"github.com/patrykjadamczyk/go-echo-log-server/config"
	"github.com/patrykjadamczyk/go-echo-log-server/logdb"
	"github.com/rakyll/statik/fs"
)

func banner(appConfiguration config.Configuration) {
	fmt.Println("Go Echo Log Admin Server")
	fmt.Println("Version:", appConfiguration.App.Version)
}

func requestLogger(targetMux http.Handler) http.Handler {
	// Handle Requests
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("ACCESS", r.URL)
		// Serve Route on HTTP
		targetMux.ServeHTTP(w, r)
	})
}

func dataRoute(appDatabase logdb.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = appDatabase.View(func(tx *logdb.Tx) error {
			html, errJson := json.Marshal(tx.GetAll())
			if errJson != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(fmt.Sprint(errJson)))
				log.Println(time.Now(), errJson)
				return errJson
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(html)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
				return err
			}
			return nil
		})
	}
}

func Main(appConfiguration config.Configuration, appDatabase logdb.DB) {
	banner(appConfiguration)
	fileName := appConfiguration.AdminServer.LogFile
	port := appConfiguration.AdminServer.Port

	fmt.Println("Preparing Log File...")
	fmt.Println("Logfile prepared. Logfile location: ", fileName)
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Println("ERROR", err)
		panic(err)
	}

	defer logFile.Close()
	// direct all log messages to log file
	log.SetOutput(logFile)

	staticFS, errStaticFS := fs.New()
	if errStaticFS != nil {
		log.Println("ERROR", errStaticFS)
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(staticFS)))
	mux.HandleFunc("/_data/", dataRoute(appDatabase))

	fmt.Println("Starting Go Logging Admin Server on port", port)
	errServer := http.ListenAndServe(fmt.Sprintf(":%s", port), requestLogger(mux))
	if errServer != nil {
		fmt.Println(errServer)
		log.Println("ERROR", errServer)
	}
}
