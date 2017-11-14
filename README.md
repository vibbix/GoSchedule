# GoSchedule
Process scheduler written in GoLang

## How to build

```bash
# make sure at least golang 1.9.2 is installed
go get github.com/vibbix/GoSchedule
cd ~/go/src/github.com/vibbix/GoSchedule
go build
./GoSchedule -h
``` 

## How to use
```
 ❯ ./GoSchedule -h
Usage of ./GoSchedule:
  -csv string
    	If location is specified, loads processes from csv file
  -fcfs
    	Run's first come first serve algorithm (default true)
  -presjf
    	Run's shortest job first algorithm premeptively (default true)
  -priority
    	Run's priority algorithm (default true)
  -rr int
    	Run's the RoundRobin Algorithm with a quantum greater than 1 (default 2)
  -sjf
    	Run's shortest job first algorithm none-premeptively (default true)
  -srtf
    	Run's ShortestRemainingTimeFirst algorithm (default true)
  -varrr int
    	Run's the variable RoundRobin Algorithm with a quantum greater than 1 (default 2)
```

## CSV Format

```csv
pid,    at, bt, priority,   period
4,      0,  3,  3,          0
3,      2,  9,  5,          0
1,      4,  3,  4,          0
5,      4,  6,  2,          0
2,      8,  6,  1,          0

```