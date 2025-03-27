//go:generate wit-bindgen-wrpc go --out-dir bindings --package github.com/Mattilsynet/map-cronjob-provider-wc/bindings wit

package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"slices"
	"strings"
	"syscall"

	"go.wasmcloud.dev/provider"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	providerCh := make(chan error, 1)
	signalCh := make(chan os.Signal, 1)
	cronHandler := New()
	p, err := provider.New(
		provider.HealthCheck(func() string {
			return handleHealthCheck(&cronHandler)
		}),
		provider.SourceLinkPut(func(link provider.InterfaceLinkDefinition) error {
			return handleNewSourceLink(&cronHandler, link)
		}),
		provider.SourceLinkDel(func(link provider.InterfaceLinkDefinition) error {
			return handleDelSourceLinks(&cronHandler, link)
		}),
		provider.Shutdown(func() error {
			return handleShutdown(&cronHandler)
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

func handleNewSourceLink(handler *Handler, link provider.InterfaceLinkDefinition) error {
	handler.provider.Logger.Info("Handling new source link", "link", link)
	if !slices.Contains(link.Interfaces, "cron-handler") {
		handler.provider.Logger.Error("Invalid source link", "error", "source link is not a cron handler")
		interfacesJoined := strings.Join(link.Interfaces, ",")
		return errors.New("the source link interfaces didn't contain 'cron-handler', got: " + interfacesJoined)
	}

	expression := link.SourceConfig["expression"]
	err := handler.AddCronJob(link.Target, expression)
	if err != nil {
		handler.provider.Logger.Error("Failed to add cron job", "error", err)
		return err
	}
	handler.StartCronjob(link.Target)
	return nil
}

func handleDelSourceLinks(handler *Handler, link provider.InterfaceLinkDefinition) error {
	handler.provider.Logger.Info("Handling del source link", "link", link)
	handler.RemoveCronJob(link.Target)
	return nil
}

func handleShutdown(handler *Handler) error {
	handler.provider.Logger.Info("Handling shutdown")
	handler.Shutdown()
	return nil
}
