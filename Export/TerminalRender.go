package Export

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/vibbix/GoSchedule/Structs"
)

// RenderToTerminal renders the current ScheduleChart to the terminal
func RenderToTerminal(sc Structs.ScheduleChart) {
	//header
	fmt.Printf("Algorithm: %v\nAverage TurnAroundTime: %v\nAverage WaitTime: %v\n", sc.AlgorithmName, sc.AverageTurnAroundTime, sc.AverageWaitTime)
	color.Set(colorHelper(int(sc.Chart[0].Process.PID)))
	fmt.Printf("%v", int(sc.Chart[0].Process.PID))
	for i := 1; i < len(sc.Chart); i++ {
		if sc.Chart[i].IsNull {
			color.Set(color.FgBlack, color.BgWhite)
			fmt.Printf("N")
			color.Set(color.BgBlack)
			continue
		} else if sc.Chart[i-1].Process.PID == sc.Chart[i].Process.PID {
			fmt.Printf(" ")
			continue
		}
		color.Set(colorHelper(int(sc.Chart[i].Process.PID)))
		fmt.Printf("%d", int(sc.Chart[i].Process.PID))
	}
	fmt.Println()
	for i := 0; i < len(sc.Chart); i++ {
		if sc.Chart[i].IsNull {
			color.Set(color.FgBlack, color.BgWhite)
			fmt.Printf("N")
			color.Set(color.BgBlack)
			continue
		}
		color.Set(colorHelper(int(sc.Chart[i].Process.PID)))
		fmt.Printf("=")
	}
}

func colorHelper(pid int) color.Attribute {
	if pid%7 == 0 {
		return color.FgMagenta
	} else if pid%6 == 0 {
		return color.FgRed
	} else if pid%5 == 0 {
		return color.FgBlue
	} else if pid%4 == 0 {
		return color.FgGreen
	} else if pid%3 == 0 {
		return color.FgYellow
	} else if pid%2 == 0 {
		return color.FgCyan
	}
	return color.FgWhite
}
