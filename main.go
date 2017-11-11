package main

import (
	"flag"

	"github.com/vibbix/GoSchedule/Parser"

	"github.com/vibbix/GoSchedule/Export"

	"github.com/vibbix/GoSchedule/Algorithms"
	"github.com/vibbix/GoSchedule/Structs"
)

var (
	csvfile     string
	runfcfs     bool
	runpriority bool
	runrr       int
	runvarrr    int
	runsjf      int
)

func main() {
	//check flags
	flag.BoolVar(&runfcfs, "fcfs", true, "Run's first come first serve algorith")
	flag.StringVar(&csvfile, "csv", "", "If location is specified, loads PID from csv file")
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
	if csvfile != "" {
		var csverr error
		processes, csverr = Parser.LoadCSV(csvfile)
		if csverr != nil {
			panic(csverr)
		}
	}
	res := Algorithms.FirstComeFirstServeSort(processes)
	// for i := 0; i < len(res); i++ {
	// 	fmt.Printf("Step %v: %+v\n", i, res[i].Process.PID)
	// }
	v := Structs.ScheduleChart{}
	v.AlgorithmName = "FirstComeFirstServe"
	v.Chart = res
	v.Processes = processes
	Export.RenderToTerminal(v)
}
