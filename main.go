package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/siebenmann/go-kstat"
	"net/http"
	"time"
)

func collectARCstats() {
	log.Debugf("Start collecting ARC stats")
	token, err := kstat.Open()
	if err != nil {
		log.Fatalf("Open failure: %s", err)
	}
	for {
		log.Debugf("Collecting...")
		ks, err := token.Lookup("zfs", 0, "arcstats")
		if err != nil {
			log.Fatalf("lookup failure on %s:0:%s: %s", "zfs", "arcstats", err)
		}
		log.Debugf("Collected: %v", ks)
		n, err := ks.GetNamed("hits")
		if err != nil {
			log.Fatalf("getting '%s' from %s: %s", "hits", ks, err)
		}
		log.Debugf("Hits: %d, type: %d, int64: %d, uint64: %d", n.UintVal, n.Type, kstat.Int64, kstat.Uint64)
		n, err = ks.GetNamed("misses")
		if err != nil {
			log.Fatalf("getting '%s' from %s: %s", "misses", ks, err)
		}
		log.Debugf("Misses: %d", n.UintVal)
		time.Sleep(10 * time.Second)
	}
}

func main() {
	log.SetLevel(log.DebugLevel)
	log.Debugf("Starting")
	go collectARCstats()
	http.Handle("/metrics", prometheus.Handler())
	err := http.ListenAndServe("0.0.0.0:9102", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
