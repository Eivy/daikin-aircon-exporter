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
	flag.StringVar(&target, "target", os.Getenv("DAIKIN_AIRCON_TARGET"), "target aircon IP Address")
	flag.BoolVar(&version, "version", false, "print version")
	flag.Parse()
	if version {
		fmt.Println(Version)
		os.Exit(0)
	}
	if target == "" {
		fmt.Fprintln(os.Stderr, "Specify target aircon IP Address with '--target' flag")
		os.Exit(1)
	}
	exporter.Run(listen, target)
}
