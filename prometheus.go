package main

import "github.com/prometheus/client_golang/prometheus"

var ticksCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "tick_recorded",
	Help: "When a new tick is recorded",
})

var buyGauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "buy_price",
	Help: "Price of buy side transactions",
})

var sellGauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "sell_price",
	Help: "Price of sell side transactions",
})
