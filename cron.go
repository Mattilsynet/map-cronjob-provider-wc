package main

import (
	"fmt"
	"strconv"
	"strings"
)

// CronParts stores the parts of a cron expression as integers
type CronParts struct {
	Seconds int
	Minutes int
	Hours   int
}

func GetCronExpressionInSeconds(cronExpr string) (int, error) {
	cronParts, err := ParseCronExpression(cronExpr)
	if err != nil {
		return 0, err
	}
	return ConvertCronToSeconds(cronParts), nil
}

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

func ConvertCronToSeconds(cronParts CronParts) int {
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
