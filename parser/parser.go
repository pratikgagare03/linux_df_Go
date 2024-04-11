package parser

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type fs struct {
	Name      string
	Size      int64
	Used      int64
	Avail     int64
	Use       int64
	MountedOn string
}

func Readable(num int) string {
	if num >= 1e6 {
		return fmt.Sprint((int64(num / 1e6))) + "G"
	} else if num >= 1e3 {
		return fmt.Sprint(num/1e3) + "M"
	} else {
		return fmt.Sprint(num) + "K"
	}

}

func PrintTop2Size(system []fs) {
	sort.Slice(system, func(i, j int) bool {
		return system[i].Size > system[j].Size
	})
	fmt.Printf("Top 2 filesystem by size are \n %+v \n %+v\n\n", system[0], system[1])
}

func PrintTop2Avail(system []fs) {
	sort.Slice(system, func(i, j int) bool {
		return system[i].Avail > system[j].Avail
	})
	fmt.Printf("Top 2 filesystem by Avail are \n %+v \n %+v\n\n", system[0], system[1])
}

func GetResults(lines []string) {
	var system []fs

	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}

		val := strings.Fields(line)

		size, _ := strconv.ParseInt(val[1], 10, 64)
		used, _ := strconv.ParseInt(val[2], 10, 64)
		avail, _ := strconv.ParseInt(val[3], 10, 64)
		use, _ := strconv.ParseInt(val[4], 10, 64)

		curr := fs{
			Name:      val[0],
			Size:      size,
			Used:      used,
			Avail:     avail,
			Use:       use,
			MountedOn: val[5],
		}
		system = append(system, curr)
	}

	PrintTop2Size(system)
	PrintTop2Avail(system)
}
