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

func init() {
	//check flags
	flag.BoolVar(&runfcfs, "fcfs", true, "Run's first come first serve algorithm")
	flag.BoolVar(&runpriority, "priority", true, "Run's priority algorithm")
	flag.StringVar(&csvfile, "csv", "", "If location is specified, loads PID from csv file")
}

func main() {
	flag.Parse()
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
	//processes := make([]Structs.Process, 0)
	if csvfile != "" {
		var csverr error
		processes, csverr = Parser.LoadCSV(csvfile)
		if csverr != nil {
			panic(csverr)
		}
	}
	algorithms := make([]Structs.ScheduleChart, 0)
	if runfcfs {
		algorithms = append(algorithms, Structs.NewScheduleChart("FirstComeFirstServe", processes, Algorithms.FirstComeFirstServeSort(processes), true))
		Export.RenderToTerminal(algorithms[len(algorithms)-1])
	}
	if runpriority {
		Structs.ResetAllProcesses(processes)
		res := Algorithms.PrioritySort(processes)
		v := Structs.ScheduleChart{}
		v.AlgorithmName = "Priority"
		v.Chart = res
		v.Processes = processes
		algorithms = append(algorithms, v)
		Export.RenderToTerminal(v)
	}
}
