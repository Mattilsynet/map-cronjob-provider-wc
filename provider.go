package main

import (
	sdk "go.wasmcloud.dev/provider"
)

// / Your Handler struct is where you can store any state or configuration that your provider needs to keep track of.
type Handler struct {
	// The provider instance
	provider *sdk.WasmcloudProvider
	// All components linked to this provider and their config.
}

// Request information about the system the provider is running on
func (h *Handler) StartCronJob() error {
	return nil
}
