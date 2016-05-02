package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

func main () {
	log.SetLevel(log.DebugLevel)
	log.Debugf("Starting")
	http.Handle("/metrics", prometheus.Handler())
	err := http.ListenAndServe("0.0.0.0:9102", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
