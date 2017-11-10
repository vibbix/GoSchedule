package Algorithms

import (
	"sort"

	"github.com/vibbix/GoSchedule/Structs"
)

// FirstComeFirstServeSort executes process's on a FCFS basis
func FirstComeFirstServeSort(processes []Structs.Process) []Structs.ProcessStep {
	//sort by arrival time, PID
	sort.Slice(processes, func(i, j int) bool {
		if processes[i].AT < processes[j].AT {
			return true
		}
		return processes[i].PID < processes[j].PID
	})
	steps := 0
	for _, p := range processes {
		steps += int(p.BT)
	}
	slices := make([]Structs.ProcessStep, steps)
	ct := 0
	for i := 0; i < len(processes); i++ {
		for j := 0; j < int(processes[i].BT); j++ {
			st := Structs.ProcessStep{Process: &processes[i], IsNull: false}
			slices[ct] = st
			ct++
		}
	}
	return slices
}
