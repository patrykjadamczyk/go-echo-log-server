package config

import "os"

type ServerConfiguration struct {
	// Port of the server
	Port string
	// File name of log file for the server
	LogFile string
}

type AppInfo struct {
	// Version of App
	Version string
	// Version Array of App
	VersionArray [4]int
}

type Configuration struct {
	// Log Server Configuration
	LogServer ServerConfiguration
	// Admin Server Configuration
	AdminServer ServerConfiguration
	// App Information
	App AppInfo
}

func FillConfigWithDefaults() Configuration {
	return Configuration{
		LogServer: ServerConfiguration{
			Port:    "7777",
			LogFile: "webrequests.log",
		},
		AdminServer: ServerConfiguration{
			Port:    "7778",
			LogFile: "error_admin.log",
		},
		App: AppInfo{
			Version:      "0.1.0.0",
			VersionArray: [4]int{0, 1, 0, 0},
		},
	}
}

func FillConfigWithEnvironmentVars() Configuration {
	// Make Config from Default Values
	newConfig := FillConfigWithDefaults()
	// Get All Environment Variables
	gelsLogServerPort, gelsLogServerPortExist := os.LookupEnv("GELS_LOG_SERVER_PORT")
	gelsLogServerLogfile, gelsLogServerLogfileExist := os.LookupEnv("GELS_LOG_SERVER_LOGFILE")
	gelsAdminServerPort, gelsAdminServerPortExist := os.LookupEnv("GELS_ADMIN_SERVER_PORT")
	gelsAdminServerLogfile, gelsAdminServerLogfileExist := os.LookupEnv("GELS_ADMIN_SERVER_LOGFILE")
	// Log Server Environment Variables Setup
	if gelsLogServerPortExist {
		newConfig.LogServer.Port = gelsLogServerPort
	}
	if gelsLogServerLogfileExist {
		newConfig.LogServer.LogFile = gelsLogServerLogfile
	}
	// Admin Server Environment Variables Setup
	if gelsAdminServerPortExist {
		newConfig.AdminServer.Port = gelsAdminServerPort
	}
	if gelsAdminServerLogfileExist {
		newConfig.AdminServer.LogFile = gelsAdminServerLogfile
	}
	// Return Changed Config
	return newConfig
}

var config = FillConfigWithDefaults()

func SetConfig(newConfig Configuration) {
	config = newConfig
}

func GetConfig() Configuration {
	return config
}

func InitConfig() Configuration {
	SetConfig(FillConfigWithEnvironmentVars())
	return GetConfig()
}
