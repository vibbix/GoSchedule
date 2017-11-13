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

// linearSort t
func linearSort(processes []Structs.Process) []Structs.ProcessStep {
	steps := 0
	for _, p := range processes {
		steps += int(p.BT)
	}
	//tempsteps := steps
	//null compensation
	// for i := 0; i < tempsteps; i++ {
	// 	notnull := false
	// 	for _, p := range processes {
	// 		//if a process is running at the current time slot, that means there is a non-null slot
	// 		if (int(p.AT) <= i) && ((int(p.AT) + int(p.BT)) >= i) {
	// 			fmt.Printf("Step #%d is between PID %d at AT %d and end point %d\n", i, p.PID, p.AT, int(p.AT)+int(p.BT))
	// 			notnull = true
	// 		}
	// 	}
	// 	if !notnull {
	// 		fmt.Printf("Incrementing at #%d\n", i)
	// 		tempsteps++
	// 	}
	// 	if i == 100 {
	// 		break
	// 	}
	// }
	slices := make([]Structs.ProcessStep, steps)
	ct := 0
	for i := 0; i < len(processes); i++ {
		// if int(processes[i].AT) < ct {
		// 	st := Structs.ProcessStep{Process: nil, IsNull: true}
		// 	slices[ct] = st
		// 	ct++
		// 	slices = append(slices, make([]Structs.ProcessStep, 1)...)
		// } else {
		for j := 0; j < int(processes[i].BT); j++ {
			if processes[i].DeIncrementBurstTime() {
				st := Structs.ProcessStep{Process: &processes[i], IsNull: false}
				slices[ct] = st
				ct++
			}
			//}
		}
	}
	return slices
}
