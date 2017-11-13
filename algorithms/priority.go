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
		if len(cproc) == 0 {
			slices[i] = Structs.ProcessStep{Process: nil, IsNull: true}
			//extend for null
			steps++
			slices = append(slices, make([]Structs.ProcessStep, 1)...)
		} else {
			//pass by value fix
			for j := 0; j < len(processes); j++ {
				if cproc[0].PID == processes[j].PID {
					if processes[j].DeIncrementBurstTime() {
						slices[i] = Structs.ProcessStep{Process: &processes[j], IsNull: false}
					}
				}
			}
			//slices[i] = Structs.ProcessStep{Process: &cproc[0], IsNull: false}
		}
	}
	return slices
}

// priorityAvilableAtTime constructs a sorted 'stack' of the available process's to handle in order
func priorityAvailableAtTime(processes []Structs.Process, time int) []Structs.Process {
	avail := []Structs.Process{}
	for _, p := range processes {
		if int(p.AT) <= time && p.GetRemainingBurstTime() > 0 {
			avail = append(avail, p)
		}
	}
	sort.Slice(avail, func(i, j int) bool {
		//return processes[i].Priority < processes[j].Priority
		if avail[i].Priority < avail[j].Priority {
			return true
		} else if avail[i].Priority == avail[j].Priority {
			if avail[i].AT < avail[j].AT {
				return true
			} else if avail[i].AT == avail[j].AT {
				return avail[i].PID < avail[j].PID
			}
		}
		return false
	})
	return avail
}
