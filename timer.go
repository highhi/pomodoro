package pomodoro

import (
	"log"
	"time"
)

// Timer start cycle.
func Timer() {
	log.Print("作業開始")

	for {
		time.Sleep(25 * time.Minute)
		log.Print("休憩")
		time.Sleep(3 * time.Minute)
		log.Print("作業再開")
	}
}
