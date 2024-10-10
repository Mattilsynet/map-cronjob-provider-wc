package main

import (
	"log"
	"time"
)

func main() {
	timer := time.NewTimer(10 * time.Second)
	stop := make(chan struct{})
	go func() {
		log.Println("starting func 1")
		log.Println("sleeping 10 seconds")
		<-timer.C
		log.Println("stopped, sleeping 1 seconds")
		<-timer.C
		stop <- struct{}{}
	}()
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("Stopping func 1")
		//		timer.Stop()
		timer.Reset(0 * time.Second)
		timer = time.NewTimer(1 * time.Second)
	}()
	<-stop
	log.Println("done")
}
