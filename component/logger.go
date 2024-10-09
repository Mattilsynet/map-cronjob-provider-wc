package main

import "github.com/Mattilsynet/map-cronjob-provider-wc/providers/logger/component/gen"

type Logger struct{}

func (lo *Logger) CronHandler() {
	component.WasiLoggingLoggingLog(component.WasiLoggingLoggingLevelInfo(), "logger", "hey we ran our cronjob")
}

func main() {}
