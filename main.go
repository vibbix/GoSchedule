package main

import (
	"fmt"

	"github.com/vibbix/GoSchedule/Export"

	"github.com/vibbix/GoSchedule/Algorithms"
	"github.com/vibbix/GoSchedule/Structs"
)

func main() {
	//check flags
	//load csv
	//parse csv
	//sanity check CSV
	//load algorithms
	//execute algorithms in parallel
	//export algorithms
	processes := []Structs.Process{
		Structs.NewProcess(1, 4, 3, 4, 0),
		Structs.NewProcess(2, 8, 6, 1, 0),
		Structs.NewProcess(3, 2, 8, 5, 0),
		Structs.NewProcess(4, 0, 3, 3, 0),
		Structs.NewProcess(5, 4, 2, 2, 0),
	}
	res := Algorithms.FirstComeFirstServeSort(processes)
	for i := 0; i < len(res); i++ {
		fmt.Printf("Step %v: %+v\n", i, res[i].Process.PID)
	}
	v := Structs.ScheduleChart{}
	v.AlgorithmName = "FirstComeFirstServe"
	v.Chart = res
	v.Processes = processes
	Export.RenderToTerminal(v)
}
