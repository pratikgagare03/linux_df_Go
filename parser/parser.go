package parser

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type fs struct {
	Name      string
	Size      string
	Used      string
	Avail     string
	Use       string
	MountedOn string
}
func GetResults(lines []string) {
	var system []fs

	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}

		val := strings.Fields(line)

		curr := fs{
			Name:      val[0],
			Size:      val[1],
			Used:      val[2],
			Avail:     val[3],
			Use:       val[4],
			MountedOn: val[5],
		}
		system = append(system, curr)
	}

	log.Writer()
	PrintTop2Size(system)
	PrintTop2Avail(system)
	PrintTop2Use(system)

}

func parseHumanReadableSize(sizeStr string) float64 {
	size,_:= strconv.ParseFloat(strings.TrimRight(sizeStr, "MKG"), 64)
	
	if(strings.Contains(sizeStr, "K")){
		size *= 1024
	}else if(strings.Contains(sizeStr, "M")){
		size *= 1024 * 1024
	}else if strings.Contains(sizeStr, "G"){
		size *= 1024 * 1024 * 1024
	}
	return size
}

func PrintTop2Size(system []fs) {
	sort.Slice(system, func(i, j int) bool {
		sizei := parseHumanReadableSize(system[i].Size)
		sizej := parseHumanReadableSize(system[j].Size)

		return sizei > sizej
	})
	log.Printf("\nTop 2 filesystem by size are :")
	fmt.Println("\nfilesystem	Size    Used    Avail   Use    MountedOn")
	for _, val:=range system[:2]{
		fmt.Print(val.Name,strings.Repeat(" ", 16-len(val.Name)),val.Size,strings.Repeat(" ", 8-len(val.Size)), val.Used, strings.Repeat(" ", 8-len(val.Used)),val.Avail, strings.Repeat(" ", 8-len(val.Avail)),val.Use, strings.Repeat(" ", 8-len(val.Use)),val.MountedOn,"\n")
	}
	fmt.Println()
}

func PrintTop2Avail(system []fs) {
	sort.Slice(system, func(i, j int) bool {
		availi := parseHumanReadableSize(system[i].Avail)
		availj := parseHumanReadableSize(system[j].Avail)
		return availi > availj
	})
	log.Printf("\nTop 2 filesystem by Avail are :")
	fmt.Println("\nName		Size    Used    Avail   Use    MountedOn")
	for _, val:=range system[:2]{
		fmt.Print(val.Name,strings.Repeat(" ", 16-len(val.Name)),val.Size,strings.Repeat(" ", 8-len(val.Size)), val.Used, strings.Repeat(" ", 8-len(val.Used)),val.Avail, strings.Repeat(" ", 8-len(val.Avail)),val.Use, strings.Repeat(" ", 8-len(val.Use)),val.MountedOn,"\n")
	}
	fmt.Println()

}

func PrintTop2Use(system []fs) {
	sort.Slice(system, func(i, j int) bool {
		usei, _ := strconv.ParseInt(strings.TrimSuffix(system[i].Use, "%"), 10, 64)
		usej, _ := strconv.ParseInt(strings.TrimSuffix(system[j].Use, "%"), 10, 64)
		return usei > usej
	})
	log.Printf("\nTop 2 filesystem by use are :")
	fmt.Println("\nName		Size    Used    Avail   Use    MountedOn")
	for _, val:=range system[:2]{
		fmt.Print(val.Name,strings.Repeat(" ", 16-len(val.Name)),val.Size,strings.Repeat(" ", 8-len(val.Size)), val.Used, strings.Repeat(" ", 8-len(val.Used)),val.Avail, strings.Repeat(" ", 8-len(val.Avail)),val.Use, strings.Repeat(" ", 8-len(val.Use)),val.MountedOn,"\n")
	}
	fmt.Println()

}

