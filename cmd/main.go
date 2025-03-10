package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Use the ticker in a loop with select
	cron := cron.New()
	l := "TZ=UTC @every 5s"
	cron.AddFunc(l, testFunc)
	go func() {
		<-sigs
		cron.Stop()
	}()
	cron.Run()
}

func testFunc() {
	log.Println("hey")
}
