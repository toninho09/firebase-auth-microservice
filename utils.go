package main

import (
	"time"
	"os"
)

func returnResponse( message string ,data ...interface{}) BaseResponse{
	return BaseResponse{Message:message,Data:data}
}

func parseTime(date string) (time.Time,error){
	return time.Parse("2006-01-02",date)
}

func getEnviroment(field string , defaultValue string) string{
	value := os.Getenv(field)
	if value == "" {
		value = defaultValue
	}
	return value
}

func getPort() string{
	return getEnviroment("PORT","8082")
}

func getEnvironmentRelease() string{
	return getEnviroment("ENV","DEV")
}