package Structs

import (
	// Used to set the data bindings for the process struct explicitly
	_ "github.com/gocarina/gocsv"
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
	PID            PID         `csv:"pid"`
	AT             ArrivalTime `csv:"at"`
	BT             BurstTime   `csv:"bt"`
	Priority       Priority    `csv:"priority"`
	Period         Period      `csv:"period"`
	remainingbt    BurstTime
	WaitTime       WaitTime
	TurnAroundTime TurnAroundTime
}

//DeIncrementBurstTime returns true if successfully deincremented
func (p *Process) DeIncrementBurstTime() bool {
	if p.remainingbt == 0 {
		return false
	}
	p.remainingbt--
	return true
}

//GetRemainingBurstTime returns the remaining amount of bursttime
func (p *Process) GetRemainingBurstTime() BurstTime {
	return p.remainingbt
}

//Reset resets the Process remaining bursttime
func (p *Process) Reset() {
	p.remainingbt++
	p.remainingbt = p.BT
}

// NewProcess intializes a new processes
func NewProcess(pid int, at int, bt int, priority int, period int) Process {
	return Process{PID: PID(pid), AT: ArrivalTime(at), BT: BurstTime(bt), Priority: Priority(priority), Period: Period(period), remainingbt: BurstTime(bt)}
}

// ProcessStep is the base unit for each step
type ProcessStep struct {
	Process *Process
	IsNull  bool
}

//ScheduleChart is base struct for rendering onto a chart
type ScheduleChart struct {
	AlgorithmName         string
	Processes             []Process
	Chart                 []ProcessStep
	AverageWaitTime       WaitTime
	AverageTurnAroundTime TurnAroundTime
}

// ComputeUtilization calculuates the TurnAroundTime and WaitTime for all processes and the efficiency
// of the scheduler
func (sc *ScheduleChart) ComputeUtilization() {
	var (
		wt  WaitTime
		tat TurnAroundTime
	)
	for i := 0; i < len(sc.Processes); i++ {
		p := sc.Processes[i]
		lastexectime := 0
		curwaittime := 0
		p.TurnAroundTime = TurnAroundTime(lastexectime - int(p.AT))
		for i := int(p.AT); i < len(sc.Chart); i++ {
			if !sc.Chart[i].IsNull && sc.Chart[i].Process.PID == p.PID {
				lastexectime = i
			}
		}
		p.TurnAroundTime = TurnAroundTime(lastexectime - int(p.AT))
		for i := int(p.AT); i < lastexectime; i++ {
			if !sc.Chart[i].IsNull && sc.Chart[i].Process.PID == p.PID {
				curwaittime++
			}
		}
		wt += WaitTime(curwaittime)
		tat += TurnAroundTime(p.TurnAroundTime)
	}
	sc.AverageWaitTime = wt / WaitTime(len(sc.Processes))
	sc.AverageTurnAroundTime = tat / TurnAroundTime(len(sc.Processes))
}

//NewScheduleChart creates a new Schedule Chart
func NewScheduleChart(name string, processes []Process, chart []ProcessStep, computeutilization bool) ScheduleChart {
	sc := ScheduleChart{AlgorithmName: name, Processes: processes, Chart: chart}
	if computeutilization {
		sc.ComputeUtilization()
	}
	return sc
}

//ResetAllProcesses resets the remaining BT in each processes
func ResetAllProcesses(processes []Process) {
	for i := 0; i < len(processes); i++ {
		processes[i].Reset()
	}
}
