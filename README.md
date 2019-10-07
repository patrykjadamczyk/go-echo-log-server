# Go Echo Log Server

## Run
1. Download Build
  * [Download the latest build from GitHub Release](https://github.com/patrykjadamczyk/go-echo-log-server/releases/latest)
  * Download from latest build on [GitHub Actions](https://github.com/patrykjadamczyk/go-echo-log-server/actions). Just go to latest successful build, in top right corner click on artifacts and download version for your system.
2. Go to your command line and run downloaded app from there
3. Admin Panel is default on port 7778 and Log Server is default on port 7777

## Configuration
You can configure some things through Environment Variables:
* `GELS_LOG_SERVER_PORT` - port of log server
* `GELS_LOG_SERVER_LOGFILE` - log file of log server
* `GELS_ADMIN_SERVER_PORT` - port of admin server
* `GELS_ADMIN_SERVER_LOGFILE` - log file of admin server
## Development
1. Go to your GOPATH and in right directory (`%GOPATH%/src/github.com/patrykjadamczyk/go-echo-log-server`) clone repo
2. `go run main.go` is starting app
