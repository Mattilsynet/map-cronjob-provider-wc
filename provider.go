package main

import (
	sdk "go.wasmcloud.dev/provider"
)

type Handler struct {
	provider *sdk.WasmcloudProvider
}

func (h *Handler) StartCronJob() error {
	for k, v := range h.provider.HostData().Config {
		h.provider.Logger.Info("Host config key: %s, value: %s", k, v)
	}
	return nil
}
