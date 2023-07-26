package logger

import (
	"fmt"
	"strconv"
	"time"
)

func Log(message string, info_type string) {
	currenTime := time.Now()
	formattedString := fmt.Sprintf("%s:%s:%s", strconv.Itoa(currenTime.Hour()), strconv.Itoa(currenTime.Minute()), strconv.Itoa(currenTime.Second()))
	switch info_type {
	case INFO:
		fmt.Printf("<%s|%s> %s\n", formattedString, info_type, message)
	case ERROR:
		fmt.Printf("<%s|\033[31m%s\033[0m> %s\n", formattedString, info_type, message)
	case SUCCESS:
		fmt.Printf("<%s|\033[32m%s\033[0m> %s\n", formattedString, info_type, message)
	case WARNING:
		fmt.Printf("<%s|\033[33m%s\033[0m> %s\n", formattedString, info_type, message)
	}
	fmt.Print("\033[0m")
}

type InfoTypes string

const (
	INFO    = "Info"
	WARNING = "Warning"
	ERROR   = "Error"
	SUCCESS = "Success"
)
