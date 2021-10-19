package scheduler

import (
	"fmt"
	"net/http"
	"github.com/jasonlvhit/gocron"
)

func Scheduler() {
	gocron.Start()
}

func ping() {
	fmt.Println("ping ping ping")
	_, _ = http.Get("https://dmba-english.herokuapp.com/")
}
