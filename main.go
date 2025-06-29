package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	webservice "github.com/BintangDwitama/quest-board-be/cmd/webservis"
)

const (
	webServiceMode = "web-service"
)

var (
	mode    string
	isLocal bool
)

func init() {
	flag.StringVar(&mode, "mode", webServiceMode, "service run mode")
	flag.Parse()
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)

	var stopFunc func()
	switch mode {
	default:
		stopFunc = webservice.Start()
	}

	for {
		<-c
		log.Println("terminate service")
		stopFunc()
		os.Exit(0)
	}
}
