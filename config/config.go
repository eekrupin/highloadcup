package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"highloadcup/travels/db"
	"highloadcup/travels/services/loggerService"
)

type HTTPServerConfig struct {
	Host         string
	InternalPort uint16
}

type AppConfig struct {
	HTTPServer *HTTPServerConfig
	DBConfig   *db.Config
	Debug      bool
	ENV        string
}

var Config AppConfig

const (
	KeyResponse = "responseData"
)

func getDefaultEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	} else {
		return value
	}
}

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		loggerService.GetMainLogger().Warn(nil, err.Error())
	}

	Config = AppConfig{
		HTTPServer: &HTTPServerConfig{},
		DBConfig:   &db.Config{},
	}

	//DB_HOST=DB_HOST;DB_PORT=1433;DB_NAME=DB_NAME;DB_USER=DB_USER;DB_PASSWORD=DB_PASSWORD;DB_MAX_OPEN_CONNS=100;DB_MAX_IDLE_CONNS=10;HTTP_INTERNAL_SERVER_PORT=8080

	Config.DBConfig.Host = os.Getenv("DB_HOST")
	DBPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic("Variable PORT from file .env must be int")
	}

	maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		panic("Variable DB_MAX_OPEN_CONNS from file .env must be int")
	}
	maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		panic("Variable DB_MAX_IDLE_CONNS from file .env must be int")
	}

	Config.DBConfig.Port = DBPort
	Config.DBConfig.DBName = os.Getenv("DB_NAME")
	Config.DBConfig.User = os.Getenv("DB_USER")
	Config.DBConfig.Password = os.Getenv("DB_PASSWORD")
	Config.DBConfig.MaxOpenConns = maxOpenConns
	Config.DBConfig.MaxIdleConns = maxIdleConns

	db.DB, err = db.Open()
	if err != nil {
		panic(err)
	}

	Config.Debug = getDefaultEnv("IS_DEBUG", "0") == "1"
	if Config.Debug {
		loggerService.GetMainLogger().Info(nil, "Environment variables")
		for _, environ := range os.Environ() {
			loggerService.GetMainLogger().Info(nil, environ)
		}
	}

	Config.ENV = getDefaultEnv("ENV", "dev")
	loggerService.GetMainLogger().Info(nil, fmt.Sprintf("ENV: %s", Config.ENV))

	Config.HTTPServer.Host = getDefaultEnv("HTTP_SERVER_HOST", "")
	httpInternalServerPort, err := strconv.ParseInt(getDefaultEnv("HTTP_INTERNAL_SERVER_PORT", "80"), 10, 32)
	if err == nil {
		Config.HTTPServer.InternalPort = uint16(httpInternalServerPort)
	} else {
		Config.HTTPServer.InternalPort = uint16(80)
	}

}
