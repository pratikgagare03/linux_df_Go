package parser

import (
	"errors"
	"project/logger"
	"strings"
)

type Fs struct {
	Name      string
	Size      string
	Used      string
	Avail     string
	Use       string
	MountedOn string
}


func GetFormattedDFOutput(system[] Fs, n int) string{
	n = min(len(system),n)
	n = max(0, n)
	ans := ""
	ans += "\nfilesystem	    Size    Used    Avail   Use    MountedOn\n"
	for _, val:=range system[:n]{
		ans += (val.Name+strings.Repeat(" ", 16-len(val.Name))+val.Size+strings.Repeat(" ", 8-len(val.Size))+ val.Used+strings.Repeat(" ", 8-len(val.Used))+val.Avail+ strings.Repeat(" ", 8-len(val.Avail))+val.Use+ strings.Repeat(" ", 8-len(val.Use))+val.MountedOn+"\n")
	}
	return ans
}

func GetResults(op string) ([]Fs, error){
	logger.Info("Parser started")
	var system []Fs
	if(len(op) == 0){
		return system,errors.New("empty input detected")
	}

	lines := strings.Split(op, "\n")
	
	
	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}

		val := strings.Fields(line)

		curr := Fs{
			Name:      val[0],
			Size:      val[1],
			Used:      val[2],
			Avail:     val[3],
			Use:       val[4],
			MountedOn: val[5],
		}
		system = append(system, curr)
	}
	logger.Info("Parser Exit")
	return system, nil
}



