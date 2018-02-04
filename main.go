package main

import "encoding/json"
import "fmt"
import "github.com/prometheus/client_golang/prometheus"
import "net/http"
import "os"

type Configuration struct {
	PGHost         string
	PGPort         int
	PGUser         string
	PGPassword     string
	PGDbname       string
	GdaxKey        string
	GdaxSecret     string
	PrometheusPort string
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: ./main <absolute path to configuration file>")
		return
	}
	file, _ := os.Open(args[0])
	decoder := json.NewDecoder(file)
	var conf = Configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error:", err)
	}

	prometheus.MustRegister(ticksCounter)
	prometheus.MustRegister(buyGauge)
	prometheus.MustRegister(sellGauge)
	http.Handle("/metrics", prometheus.Handler())
	go http.ListenAndServe(conf.PrometheusPort, nil)

	ingester := NewIngester(&conf)
	ingester.Start()
}
