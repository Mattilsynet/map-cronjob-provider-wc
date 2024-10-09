//go:generate wit-bindgen-wrpc go --out-dir bindings --package github.com/Mattilsynet/map-cronjob-provider-wc/bindings wit

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.wasmcloud.dev/provider"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Initialize the provider with callbacks to track linked components
	cronHandler := Handler{}
	p, err := provider.New()
	if err != nil {
		return err
	}
	cronHandler.provider = p

	providerCh := make(chan error, 1)
	signalCh := make(chan os.Signal, 1)

	go func() {
		err := cronHandler.StartCronJob()
		providerCh <- err
	}()
	signal.Notify(signalCh, syscall.SIGINT)

	<-signalCh
	p.Shutdown()

	return nil
}
