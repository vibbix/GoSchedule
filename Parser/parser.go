package Parser

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/vibbix/GoSchedule/Structs"
)

// LoadCSV loads the CSV into a slice of process's
func LoadCSV(uri string) ([]Structs.Process, error) {
	var (
		processlist []Structs.Process
	)
	file, fileerr := os.Open(uri)
	if fileerr != nil {
		return nil, fileerr
	}
	defer file.Close()
	csverr := gocsv.Unmarshal(file, &processlist)
	if csverr != nil {
		return nil, csverr
	}
	return processlist, nil
}
