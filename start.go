package main

import (
	"flag"
	"github.com/dmba-english/controller"
	. "github.com/dmba-english/scheduler"
	"github.com/dmba-english/telegram"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	Scheduler()
	telegram.Tg()
	port := flag.String("port", os.Getenv("PORT"), "app port")
	if len(*port) == 0 {
		*port = "3000"
	}
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", controller.PingHandler)
	http.Handle("/", rtr)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
