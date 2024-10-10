//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world component --out gen ./wit
package main

import (
	"github.com/Mattilsynet/map-cronjob-provider-wc/providers/logger/component/gen/mattilsynet/cronjob/cronjob"
	"go.wasmcloud.dev/component/log/wasilog"
)

func init() {
	cronjob.Exports.CronHandler = func() {
		logger := wasilog.ContextLogger("cronjob-handler")
		logger.Info("Cronjob handler called")
	}
}

func main() {
}
