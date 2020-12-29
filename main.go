package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Eivy/daikin-aircon-exporter/exporter"
)

func main() {
	var listen string
	var target string
	var version bool
	flag.StringVar(&listen, "listen address", "0.0.0.0:9823", "listen address")
	flag.StringVar(&target, "target", "", "target aircon IP")
	flag.BoolVar(&version, "version", false, "print version")
	flag.Parse()
	if version {
		fmt.Println(Version)
		os.Exit(0)
	}
	exporter.Run(listen, target)
}
