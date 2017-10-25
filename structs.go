package main

import (
	_ "encoding/csv"
)

//PID is the Process ID
type PID int

// BurstTime is the time the process requests when it comes into the scheduler
type BurstTime int

// ArrivalTime is the time at which the process arrives at the scheduler
type ArrivalTime int

// Priority is the
type Priority int

// Process is the base struct for storing process info
type Process struct {
	PID      PID         `csv:"pid"`
	BT       BurstTime   `csv:"bt"`
	AT       ArrivalTime `csv:"at"`
	Priority Priority    `csv:"priority"`
}

// ProcessStep is the base unit for each step
type ProcessStep struct {
	Process *Process
	IsNull  bool
}
