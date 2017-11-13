// Package Algorithms is the class containing implementations of all the
// CPU scheduling algorithms
package Algorithms

import (
	"github.com/vibbix/GoSchedule/Structs"
)

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
