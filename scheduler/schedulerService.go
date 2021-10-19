package scheduler

import (
	"github.com/jasonlvhit/gocron"
	"log"
	"net/http"
)

func Scheduler() {
	gocron.Every(10).Minutes().Do(ping)
	gocron.Start()
}

func ping() {
	log.Println("ping ping ping")
	_, _ = http.Get("https://dmba-english.herokuapp.com/")
}
