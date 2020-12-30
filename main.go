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
	port := os.Getenv("DAIKIN_AIRCON_PORT")
	if port == "" {
		port = "9823"
	}
	flag.StringVar(&listen, "listen port", port, "listen address")
	flag.StringVar(&target, "target", os.Getenv("DAIKIN_AIRCON_TARGET"), "target aircon IP Address")
	flag.BoolVar(&version, "version", false, "print version")
	flag.Parse()
	if version {
		fmt.Println(exporter.Version)
		os.Exit(0)
	}
	if target == "" {
		fmt.Fprintln(os.Stderr, "Specify target aircon IP Address with '--target' flag")
		os.Exit(1)
	}
	exporter.Run(listen, target)
}
