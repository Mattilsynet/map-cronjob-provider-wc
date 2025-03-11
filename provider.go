package main

import (
	"context"

	cronJobComponent "github.com/Mattilsynet/map-cronjob-provider-wc/bindings/mattilsynet/cronjob/cronjob"
	"github.com/robfig/cron/v3"
	sdk "go.wasmcloud.dev/provider"
)

type Handler struct {
	provider *sdk.WasmcloudProvider
	linkedTo map[string]map[string]string
	cron     map[string]*cron.Cron
}

func New() Handler {
	return Handler{
		linkedTo: make(map[string]map[string]string),
		cron:     make(map[string]*cron.Cron),
	}
}

func (h *Handler) AddCronJob(target string, expression string) error {
	cron := cron.New()
	client := h.provider.OutgoingRpcClient(target)
	_, err := cron.AddFunc(expression,
		func() {
			cronJobComponent.CronHandler(context.TODO(), client)
		},
	)
	if err != nil {
		return err
	}
	h.cron[target] = cron
	return nil
}

func (h *Handler) StartCronjob(target string) {
	h.cron[target].Start()
}

func (h *Handler) StopCronJob(target string) {
	h.cron[target].Stop()
}

func (h *Handler) RemoveCronJob(target string) {
	h.cron[target] = nil
}

func (h *Handler) Shutdown() {
	for target := range h.cron {
		h.StopCronJob(target)
	}
}
