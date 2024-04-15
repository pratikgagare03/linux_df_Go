package parser

import (
	"sort"
	"strconv"
	"strings"
	"project/logger"
)

type fs struct {
	Name      string
	Size      string
	Used      string
	Avail     string
	Use       string
	MountedOn string
}


func GetFormattedDFOutput(system[] fs, n int) string{
	n = min(len(system),n)
	n = max(0, n)
	ans := ""
	ans += "\nfilesystem	    Size    Used    Avail   Use    MountedOn\n"
	for _, val:=range system[:n]{
		ans += (val.Name+strings.Repeat(" ", 16-len(val.Name))+val.Size+strings.Repeat(" ", 8-len(val.Size))+ val.Used+strings.Repeat(" ", 8-len(val.Used))+val.Avail+ strings.Repeat(" ", 8-len(val.Avail))+val.Use+ strings.Repeat(" ", 8-len(val.Use))+val.MountedOn+"\n")
	}
	return ans
}

func GetResults(op string) string{
	logger.Info("Parser started")
	lines := strings.Split(op, "\n")

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
	res := GetTop2Size(system)
	if(len(res)!=0){
		logger.Debug("GetTop2Size returned a non zero length string")
	}
	res1 := GetTop2Avail(system)
	if(len(res1)!=0){
		logger.Debug("GetTop2Avail returned a non zero length string")
		res+=res1
	}
	res2 := GetTop2Use(system)
	if(len(res2)!=0){
		logger.Debug("GetTop2Use returned a non zero length string")
		res+=res2
	}
	logger.Info("Parser exit")
	return res
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

func GetTop2Size(system []fs) string{
	sort.Slice(system, func(i, j int) bool {
		sizei := parseHumanReadableSize(system[i].Size)
		sizej := parseHumanReadableSize(system[j].Size)
		return sizei > sizej
	})
	return "\nTop 2 filesystem by size are :" + GetFormattedDFOutput(system, 2)
}

func GetTop2Avail(system []fs) string{
	sort.Slice(system, func(i, j int) bool {
		availi := parseHumanReadableSize(system[i].Avail)
		availj := parseHumanReadableSize(system[j].Avail)
		return availi > availj
	})
	return "\nTop 2 filesystem by Avail are :" + GetFormattedDFOutput(system, 2)
}

func GetTop2Use(system []fs) string{
	sort.Slice(system, func(i, j int) bool {
		usei, _ := strconv.ParseInt(strings.TrimSuffix(system[i].Use, "%"), 10, 64)
		usej, _ := strconv.ParseInt(strings.TrimSuffix(system[j].Use, "%"), 10, 64)
		return usei > usej
	})
	return "\nTop 2 filesystem by use are :" + GetFormattedDFOutput(system, 2)
}

