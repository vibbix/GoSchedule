package Algorithms

import (
	"sort"

	"github.com/vibbix/GoSchedule/Structs"
)

//RoundRobinSort uses the RoundRobin algorithm to schedule processes
func RoundRobinSort(processes []Structs.Process, quantum int, variable bool) []Structs.ProcessStep {
	lastexectime := make([]int, len(processes))
	steps := 0
	for _, proc := range processes {
		steps += int(proc.BT)
	}
	slices := make([]Structs.ProcessStep, steps)
	for i := 0; i < steps; i++ {
		cproc := getRRStack(processes, i, lastexectime)
		if i > 1 && slices[i-1].IsNull == false && slices[i-1].Process == nil {
			i -= 2
			continue
		}
		if len(cproc) == 0 {
			slices[i] = Structs.ProcessStep{Process: nil, IsNull: true}
			steps++
			slices = append(slices, make([]Structs.ProcessStep, 1)...)
		} else {
			cp, cpi := getProcess(processes, int(cproc[0].PID))
			for j := 0; j < quantum; j++ {
				if cp.DeIncrementBurstTime() {
					slices[i] = Structs.ProcessStep{Process: cp, IsNull: false}
					lastexectime[cpi] = i
					//i++
				} else if !variable {
					slices[i] = Structs.ProcessStep{Process: nil, IsNull: true}
					steps++
					slices = append(slices, make([]Structs.ProcessStep, 1)...)
					//i++
				} else {
					break
				}
				i++
			}
		}
	}
	return slices
}

func getRRStack(processes []Structs.Process, time int, lastexectime []int) []Structs.Process {
	avail := []Structs.Process{}
	for _, p := range processes {
		if int(p.AT) <= time && p.GetRemainingBurstTime() > 0 {
			avail = append(avail, p)
		}
	}
	sort.Slice(avail, func(i, j int) bool {
		if lastexectime[i] < lastexectime[j] {
			return true
		} else if lastexectime[i] == lastexectime[j] {
			return avail[i].AT < avail[j].AT
		}
		return false
	})
	return avail
}
