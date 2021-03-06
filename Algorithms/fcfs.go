package Algorithms

import (
	"sort"

	"github.com/vibbix/GoSchedule/Structs"
)

// FirstComeFirstServeSort executes process's on a FCFS basis
func FirstComeFirstServeSort(processes []Structs.Process) []Structs.ProcessStep {
	steps := 0
	for _, proc := range processes {
		steps += int(proc.BT)
	}
	slices := make([]Structs.ProcessStep, steps)
	for i := 0; i < steps; i++ {
		cproc := getFCFSstack(processes, i)
		if len(cproc) == 0 {
			slices[i] = Structs.ProcessStep{Process: nil, IsNull: true}
			steps++
			slices = append(slices, make([]Structs.ProcessStep, 1)...)
		} else {
			//pass by value fix
			cp, _ := getProcess(processes, int(cproc[0].PID))
			if cp.DeIncrementBurstTime() {
				slices[i] = Structs.ProcessStep{Process: cp, IsNull: false}
			}
			//slices[i] = Structs.ProcessStep{Process: &cproc[0], IsNull: false}
		}
	}
	return slices
}

func getFCFSstack(processes []Structs.Process, time int) []Structs.Process {
	avail := []Structs.Process{}
	for _, p := range processes {
		if int(p.AT) <= time && p.GetRemainingBurstTime() > 0 {
			avail = append(avail, p)
		}
	}
	sort.Slice(avail, func(i, j int) bool {
		//return processes[i].Priority < processes[j].Priority
		if avail[i].AT < avail[j].AT {
			return true
		} else if avail[i].Priority == avail[j].Priority {
			return avail[i].PID < avail[j].PID
		}
		return false
	})
	return avail
}
