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
	providerCh := make(chan error, 1)
	signalCh := make(chan os.Signal, 1)
	cronHandler := New()
	p, err := provider.New(
		provider.HealthCheck(func() string {
			return handleHealthCheck(&cronHandler)
		}),
		provider.SourceLinkPut(func(link provider.InterfaceLinkDefinition) error {
			return handleNewSourceLink(&cronHandler, link, signalCh)
		}),
		provider.SourceLinkDel(func(link provider.InterfaceLinkDefinition) error {
			return handleDelSourceLinks(&cronHandler, link)
		}),
	)
	if err != nil {
		return err
	}
	cronHandler.provider = p
	go func() {
		err := p.Start()
		providerCh <- err
	}()
	go func() {
	}()
	signal.Notify(signalCh, syscall.SIGINT)

	<-signalCh
	p.Shutdown()

	return nil
}

func handleHealthCheck(_ *Handler) string {
	return "provider healthy"
}

func handleNewSourceLink(handler *Handler, link provider.InterfaceLinkDefinition, signalCh chan os.Signal) error {
	handler.provider.Logger.Info("Handling new source link", "link", link)
	handler.StartCronJob(signalCh, link.Target, link.TargetConfig["expression"])
	handler.linkedTo[link.Target] = link.SourceConfig
	return nil
}

func handleDelSourceLinks(handler *Handler, link provider.InterfaceLinkDefinition) error {
	handler.provider.Logger.Info("Handling del source link", "link", link)
	handler.RemoveCronJobs()
	delete(handler.linkedTo, link.SourceID)
	return nil
}

func handleShutdown(handler *Handler) error {
	handler.provider.Logger.Info("Handling shutdown")
	clear(handler.linkedTo)
	return nil
}
