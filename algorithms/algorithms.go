// Package Algorithms is the class containing implementations of all the
// CPU scheduling algorithms
package Algorithms

import (
	"github.com/vibbix/GoSchedule/Structs"
)

// ComputeUtilization calculuates the TurnAroundTime and WaitTime for all processes and the efficiency
// of the scheduler
func ComputeUtilization(sc Structs.ScheduleChart) {
	var (
		wt  Structs.WaitTime
		tat Structs.TurnAroundTime
	)
	for _, p := range sc.Processes {
		lastexectime := 0
		curwaittime := 0
		p.TurnAroundTime = Structs.TurnAroundTime(lastexectime - int(p.AT))
		for i := int(p.AT); i < len(sc.Chart); i++ {
			if !sc.Chart[i].IsNull && sc.Chart[i].Process.PID == p.PID {
				lastexectime = i
			}
		}
		p.TurnAroundTime = Structs.TurnAroundTime(lastexectime - int(p.AT))
		for i := int(p.AT); i < lastexectime; i++ {
			if !sc.Chart[i].IsNull && sc.Chart[i].Process.PID == p.PID {
				curwaittime++
			}
		}
		wt += Structs.WaitTime(curwaittime)
		tat += Structs.TurnAroundTime(p.TurnAroundTime)
	}
	sc.AverageWaitTime = wt / Structs.WaitTime(len(sc.Processes))
	sc.AverageTurnAroundTime = tat / Structs.TurnAroundTime(len(sc.Processes))
}
