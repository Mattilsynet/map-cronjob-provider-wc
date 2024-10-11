package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// CronParts stores the parts of a cron expression as integers
type CronParts struct {
	Seconds int
	Minutes int
	Hours   int
}

// ConvertCronExpressionToTicker creates a time.Ticker based on the cron expression
func ConvertCronExpressionToTicker(cronExpr string) (*time.Ticker, error) {
	cronParts, err := ParseCronExpression(cronExpr)
	if err != nil {
		return nil, err
	}

	// Convert cron parts into duration
	duration := ConvertCronToDuration(cronParts)
	if duration <= 0 {
		return nil, fmt.Errorf("invalid duration")
	}

	// Create and return a ticker with the calculated duration
	return time.NewTicker(duration), nil
}

// ParseCronExpression parses a cron expression into CronParts
func ParseCronExpression(cronExpr string) (CronParts, error) {
	parts := strings.Split(cronExpr, " ")
	if len(parts) != 3 {
		return CronParts{}, fmt.Errorf("invalid cron expression format")
	}

	// Parse each part of the cron expression, handling "*" as a wildcard (-1)
	seconds := parsePart(parts[0])
	minutes := parsePart(parts[1])
	hours := parsePart(parts[2])

	return CronParts{
		Seconds: seconds,
		Minutes: minutes,
		Hours:   hours,
	}, nil
}

func parsePart(part string) int {
	if part == "*" {
		return -1 // Treat "*" as "every"
	}
	// Convert string to int safely since we don't need error handling for this task
	val, _ := strconv.Atoi(part)
	return val
}

// ConvertCronToDuration converts the cron parts into a time.Duration
func ConvertCronToDuration(cronParts CronParts) time.Duration {
	totalSeconds := 0

	// Handle seconds
	if cronParts.Seconds != -1 {
		totalSeconds += cronParts.Seconds
	} else {
		return time.Second // "*" for seconds, tick every second
	}

	// Handle minutes
	if cronParts.Minutes != -1 {
		totalSeconds += cronParts.Minutes * 60
	} else if cronParts.Hours == -1 { // if hours is also "*", tick every minute
		return time.Minute
	}

	// Handle hours
	if cronParts.Hours != -1 {
		totalSeconds += cronParts.Hours * 3600
	}

	// Convert total seconds to time.Duration
	return time.Duration(totalSeconds) * time.Second
}

func main() {
	// Example: Parse cron expression and create a ticker
	cronExpr := "5 * *" // Every 5 seconds
	ticker, err := ConvertCronExpressionToTicker(cronExpr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer ticker.Stop() // Stop the ticker when done

	// Capture OS signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Use the ticker in a loop with select
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Ticker ticked at:", t)

		case sig := <-sigs:
			fmt.Println("Received signal:", sig)
			fmt.Println("Shutting down gracefully...")
			return
		}
	}
}
