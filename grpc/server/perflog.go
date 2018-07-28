package main

import "time"

// PerformanceTrace records the timing points during DB operations
type PerformanceTrace struct {
	RequestID    string
	TotalElapsed time.Duration
	Start        time.Time
	End          time.Time
	Logs         []*PerformanceLog
}

// PerformanceLog
type PerformanceLog struct {
}
