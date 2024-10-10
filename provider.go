package main

import (
	"context"
	"os"
	"time"

	"github.com/Mattilsynet/map-cronjob-provider-wc/bindings/mattilsynet/cronjob/cronjob"
	sdk "go.wasmcloud.dev/provider"
)

type CronJob struct {
	target         string
	cronExpression string
	lastRun        time.Time
}
type Handler struct {
	provider      *sdk.WasmcloudProvider
	linkedTo      map[string]map[string]string
	sleepDuration int
	cronJobs      map[string]*CronJob
	timer         *time.Timer
}

func (h *Handler) StartCronJob(signal chan os.Signal) {
	log := h.provider.Logger
outerloop:
	for {
		select {
		case <-signal:
			break outerloop
		case <-h.timer.C:
			if h.shouldRun() {
			}
			h.sleepDuration = h.ConfigureNextRunningCronJob()
			h.timer = time.NewTimer(time.Duration(h.sleepDuration) * time.Second)
		}
	}
}

func (h *Handler) ConfigureNextRunningCronJob() int {
	h.timer.Reset(0 * time.Second)
}

func (h *Handler) shouldRun() bool {
}
