package Algorithms

import (
	"sort"

	"github.com/vibbix/GoSchedule/Structs"
)

// PrioritySort sorts the process's by their priorty to execute
func PrioritySort(processes []Structs.Process) []Structs.ProcessStep {
	sort.Slice(processes, func(i, j int) bool {
		if processes[i].AT < processes[j].AT {
			return true
		} else if processes[i].Priority < processes[j].Priority {
			return true
		}
		return processes[i].PID < processes[j].PID
	})
	return linearSort(processes)
}
