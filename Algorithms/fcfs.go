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

	return linearSort(processes)
}
