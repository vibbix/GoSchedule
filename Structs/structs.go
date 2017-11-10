package Structs

import (
	// Used to set the data bindings for the process struct explicitly
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

// Period is used in Rate Monotonic Scheduling & Earliest Deadline First to determine
// How frequently a process is run
type Period int

// WaitTime is the amount of time the process
type WaitTime int

// TurnAroundTime is the amount of time it takes for a program to complete once its arrived
type TurnAroundTime int

// Process is the base struct for storing process info
type Process struct {
	PID      PID         `csv:"pid"`
	AT       ArrivalTime `csv:"at"`
	BT       BurstTime   `csv:"bt"`
	Priority Priority    `csv:"priority"`
	Period   Period      `csv:"period"`
}

// ProcessTable is attached to each process
type ProcessTable struct {
	WT  WaitTime
	TAT TurnAroundTime
}

// ProcessStep is the base unit for each step
type ProcessStep struct {
	Process *Process
	IsNull  bool
}

//ScheduleChart is base struct for rendering onto a chart
type ScheduleChart struct {
	AlgorithmName         string
	Chart                 []ProcessStep
	AverageWaitTime       WaitTime
	AverageTurnAroundTime TurnAroundTime
}