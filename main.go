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
	cronHandler := Handler{
		linkedTo: make(map[string]map[string]string),
	}
	p, err := provider.New(
		provider.HealthCheck(func() string {
			return handleHealthCheck(&cronHandler)
		}),
		provider.SourceLinkPut(func(link provider.InterfaceLinkDefinition) error {
			return handleNewSourceLink(&cronHandler, link)
		}),
		provider.SourceLinkDel(func(link provider.InterfaceLinkDefinition) error {
			return handleDelSourceLink(&cronHandler, link)
		}),
	)
	if err != nil {
		return err
	}
	cronHandler.provider = p
	providerCh := make(chan error, 1)
	signalCh := make(chan os.Signal, 1)
	go func() {
		err := p.Start()
		providerCh <- err
	}()
	go func() {
		cronHandler.StartCronJob(signalCh)
	}()
	signal.Notify(signalCh, syscall.SIGINT)

	<-signalCh
	p.Shutdown()

	return nil
}

func handleHealthCheck(_ *Handler) string {
	return "provider healthy"
}

func handleNewSourceLink(handler *Handler, link provider.InterfaceLinkDefinition) error {
	handler.provider.Logger.Info("Handling new source link", "link", link)
	handler.linkedTo[link.Target] = link.SourceConfig
	return nil
}

func handleDelSourceLink(handler *Handler, link provider.InterfaceLinkDefinition) error {
	handler.provider.Logger.Info("Handling del source link", "link", link)
	delete(handler.linkedTo, link.SourceID)
	return nil
}

func handleShutdown(handler *Handler) error {
	handler.provider.Logger.Info("Handling shutdown")
	clear(handler.linkedTo)
	return nil
}
