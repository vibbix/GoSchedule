package Algorithms

import (
	"sort"

	"github.com/vibbix/GoSchedule/Structs"
)

//NonePreemptiveShortestJobFirstSort sorts using a SRTF approach, none-preemptively
func NonePreemptiveShortestJobFirstSort(processes []Structs.Process) []Structs.ProcessStep {
	steps := 0
	for _, proc := range processes {
		steps += int(proc.BT)
	}
	slices := make([]Structs.ProcessStep, steps)
	for i := 0; i < steps; i++ {
		cproc := srtfGetStack(processes, i)
		if len(cproc) == 0 {
			slices[i] = Structs.ProcessStep{Process: nil, IsNull: true}
			//extend for null
			steps++
			slices = append(slices, make([]Structs.ProcessStep, 1)...)
		} else {
			cp, _ := getProcess(processes, int(cproc[0].PID))
			for cp.DeIncrementBurstTime() {
				//fmt.Printf("Step #%d: PID %d, rBT %d\n", i, cp.PID, cp.GetRemainingBurstTime())
				slices[i] = Structs.ProcessStep{Process: cp, IsNull: false}
				i++
			}
			i--
		}
	}
	return slices
}

//PreemptiveShortestJobFirstSort sorts using a SRTF approach, preemptively
func PreemptiveShortestJobFirstSort(processes []Structs.Process) []Structs.ProcessStep {
	steps := 0
	for _, proc := range processes {
		steps += int(proc.BT)
	}
	slices := make([]Structs.ProcessStep, steps)
	for i := 0; i < steps; i++ {
		cproc := srtfGetStack(processes, i)
		if len(cproc) == 0 {
			slices[i] = Structs.ProcessStep{Process: nil, IsNull: true}
			//extend for null
			steps++
			slices = append(slices, make([]Structs.ProcessStep, 1)...)
		} else {
			//pass by value fix
			cp, _ := getProcess(processes, int(cproc[0].PID))
			if cp.DeIncrementBurstTime() {
				slices[i] = Structs.ProcessStep{Process: cp, IsNull: false}
			}
		}
	}
	return slices
}

// srtfGetStack gets the upcoming list of processes to sort
func preemptiveSJFGetStack(processes []Structs.Process, time int) []Structs.Process {
	avail := []Structs.Process{}
	for _, p := range processes {
		if int(p.AT) <= time && p.GetRemainingBurstTime() > 0 {
			avail = append(avail, p)
		}
	}
	sort.Slice(avail, func(i, j int) bool {
		if avail[i].BT < avail[j].BT {
			return true
		} else if avail[i].BT == avail[j].BT {
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
