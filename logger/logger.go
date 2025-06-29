package logger

import (
	"fmt"
	"log"
	"time"
)

func Log(logRestrinction bool, logStr string) {
	if logRestrinction {
		t := time.Now()
		formattedTimestamp := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		log.Println(formattedTimestamp, " ", logStr)
	}
}

func Fatal(message string, err error) {
	log.Fatal(message, err)
}
