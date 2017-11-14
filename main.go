package main

import (
	"flag"
	"fmt"

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
	runsjf      bool
	runpresjf   bool
	runsrtf     bool
)

func init() {
	//check flags
	flag.BoolVar(&runfcfs, "fcfs", true, "Run's first come first serve algorithm")
	flag.BoolVar(&runsjf, "sjf", true, "Run's shortest job first algorithm none-premeptively")
	flag.BoolVar(&runpresjf, "presjf", true, "Run's shortest job first algorithm premeptively")
	flag.BoolVar(&runpriority, "priority", true, "Run's priority algorithm")
	flag.BoolVar(&runsrtf, "srtf", true, "Run's ShortestRemainingTimeFirst algorithm")
	flag.StringVar(&csvfile, "csv", "/Users/vibbix/go/src/github.com/vibbix/GoSchedule/ex2.csv", "If location is specified, loads PID from csv file")
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
	} else {
		fmt.Println("Using default values")
	}
	Export.PrintTable(processes)
	algorithms := make([]Structs.ScheduleChart, 0)
	if runfcfs {
		Structs.ResetAllProcesses(processes)
		algorithms = append(algorithms, Structs.NewScheduleChart("FirstComeFirstServe", processes, Algorithms.FirstComeFirstServeSort(processes), true))
		Export.RenderToTerminal(algorithms[len(algorithms)-1])
	}
	if runpriority {
		Structs.ResetAllProcesses(processes)
		algorithms = append(algorithms, Structs.NewScheduleChart("Priority", processes, Algorithms.PrioritySort(processes), true))
		Export.RenderToTerminal(algorithms[len(algorithms)-1])
	}
	if runsrtf {
		Structs.ResetAllProcesses(processes)
		algorithms = append(algorithms, Structs.NewScheduleChart("ShortestRemainingTimeFirst", processes, Algorithms.ShortestRemainingTimeFirstSort(processes), true))
		Export.RenderToTerminal(algorithms[len(algorithms)-1])
	}
	if runpresjf {
		Structs.ResetAllProcesses(processes)
		algorithms = append(algorithms, Structs.NewScheduleChart("PreemptiveShortestJobFirst", processes, Algorithms.PreemptiveShortestJobFirstSort(processes), true))
		Export.RenderToTerminal(algorithms[len(algorithms)-1])
	}
	if runsjf {
		Structs.ResetAllProcesses(processes)
		algorithms = append(algorithms, Structs.NewScheduleChart("ShortestJobFirst", processes, Algorithms.NonePreemptiveShortestJobFirstSort(processes), true))
		Export.RenderToTerminal(algorithms[len(algorithms)-1])
	}
}
