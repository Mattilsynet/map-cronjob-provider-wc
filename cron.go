package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// CronParts stores the parts of a cron expression as integers
type CronParts struct {
	Seconds int
	Minutes int
	Hours   int
}

func ConvertToTicker(cronExpr string) (*time.Ticker, error) {
	if cronExpr == "" {
		return nil, errors.New("expression cannot be empty in cron source config")
	}
	cronParts, err := ParseCronExpression(cronExpr)
	if err != nil {
		return nil, err
	}
	seconds := ConvertCronPartsToSeconds(cronParts)
	duration := time.Duration(seconds) * time.Second
	return time.NewTicker(duration), nil
}

func ParseCronExpression(cronExpr string) (CronParts, error) {
	parts := strings.Split(cronExpr, " ")
	if len(parts) != 3 {
		return CronParts{}, fmt.Errorf("invalid cron expression format")
	}
	// Parse each part of the cron expression, handling "*" as a wildcard (-1)
	seconds, err := parsePart(parts[0])
	if err != nil {
		return CronParts{}, err
	}
	minutes, err := parsePart(parts[1])
	if err != nil {
		return CronParts{}, err
	}
	hours, err := parsePart(parts[2])
	if err != nil {
		return CronParts{}, err
	}

	return CronParts{
		Seconds: seconds,
		Minutes: minutes,
		Hours:   hours,
	}, nil
}

func parsePart(part string) (int, error) {
	if part == "*" {
		return -1, nil // Treat "*" as "every"
	}
	// Convert string to int safely since we don't need error handling for this task
	val, err := strconv.Atoi(part)
	return val, err
}

func ConvertCronPartsToSeconds(cronParts CronParts) int {
	totalSeconds := 0

	if cronParts.Seconds != -1 {
		totalSeconds += cronParts.Seconds
	}
	if cronParts.Minutes != -1 {
		totalSeconds += cronParts.Minutes * 60
	}
	if cronParts.Hours != -1 {
		totalSeconds += cronParts.Hours * 3600
	}
	return totalSeconds
}
