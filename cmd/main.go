package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"
)

func main() {
	// Example: Parse cron expression and create a ticker
	// cronExpr := "5 * *" // Every 5 seconds
	// ticker, err := ConvertCronExpressionToTicker(cronExpr)

	// Capture OS signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Use the ticker in a loop with select
	for {
		select {
		case sig := <-sigs:
			fmt.Println("Received signal:", sig)
			fmt.Println("Shutting down gracefully...")
			return
		}
	}
	cron.New()
}
