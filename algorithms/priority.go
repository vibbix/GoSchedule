package Algorithms

import (
	"sort"

	"github.com/vibbix/GoSchedule/Structs"
)

// PrioritySort sorts the process's by their priorty to execute
func PrioritySort(processes []Structs.Process) []Structs.ProcessStep {
	steps := 0

	for _, proc := range processes {
		steps += int(proc.BT)
	}
	slices := make([]Structs.ProcessStep, steps)
	for i := 0; i < steps; i++ {
		cproc := priorityAvailableAtTime(processes, i)
		if len(cproc) == 0 || !cproc[0].DeIncrementBurstTime() {
			slices[i] = Structs.ProcessStep{Process: nil, IsNull: true}
			//extend for null
			steps++
			slices = append(slices, make([]Structs.ProcessStep, 1)...)
		} else {
			slices[i] = Structs.ProcessStep{Process: &cproc[0], IsNull: false}
		}
	}
	return slices
}

func priorityAvailableAtTime(processes []Structs.Process, time int) []Structs.Process {
	avail := []Structs.Process{}
	for _, p := range processes {
		if int(p.AT) <= time && p.GetRemainingBurstTime() > 0 {
			avail = append(avail, p)
		}
	}
	sort.Slice(avail, func(i, j int) bool {
		if processes[i].Priority > processes[j].Priority {
			return true
		} else if processes[i].AT > processes[j].AT {
			return true
		}
		return processes[i].PID < processes[j].PID
	})
	return avail
}
