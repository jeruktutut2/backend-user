package util

import (
	"log"
	"time"
)

func SetTimezone(timezone string) {
	location, err := time.LoadLocation(timezone)
	time.Local = location
	if err != nil {
		log.Fatal(err)
	}
}
