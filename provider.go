package main

import (
	"context"
	"os"
	"time"

	cronJobComponent "github.com/Mattilsynet/map-cronjob-provider-wc/bindings/mattilsynet/cronjob/cronjob"
	sdk "go.wasmcloud.dev/provider"
)

type CronJob struct {
	target       string
	ticker       *time.Ticker
	removeSignal chan struct{}
}
type Handler struct {
	provider *sdk.WasmcloudProvider
	linkedTo map[string]map[string]string
	cronJobs map[string]*CronJob
}

func New() Handler {
	return Handler{
		linkedTo: make(map[string]map[string]string),
		cronJobs: make(map[string]*CronJob),
	}
}

func (h *Handler) StartCronJob(osSignal <-chan os.Signal, target string, expression string) error {
	ticker, err := ConvertToTicker(expression)
	if err != nil {
		return err
	}
	cronjob := &CronJob{
		target:       target,
		removeSignal: make(chan struct{}),
		ticker:       ticker,
	}
	h.cronJobs[target] = cronjob
	go func(osSignal <-chan os.Signal, cronjob *CronJob) {
		for {
			select {
			case <-osSignal:
				return
			case <-cronjob.removeSignal:
				return
			case <-cronjob.ticker.C:
				client := h.provider.OutgoingRpcClient(cronjob.target)
				cronJobComponent.CronHandler(context.TODO(), client)
			}
		}
	}(osSignal, cronjob)
	return nil
}

func (h *Handler) RemoveCronJobs(target string) {
	for key, cronJob := range h.cronJobs {
		if cronJob.target == target {
			close(cronJob.removeSignal)
			delete(h.cronJobs, key)
		}
	}
}
