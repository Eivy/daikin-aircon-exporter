package main

import (
	"flag"

	"github.com/eivy/daikin-aircon-exporter/exporter"
)

func main() {
	var listen string
	var target string
	flag.StringVar(&listen, "listen address", "0.0.0.0:9823", "listen address")
	flag.StringVar(&target, "target", "", "target aircon IP")
	flag.Parse()
	exporter.Run(listen, target)
}
