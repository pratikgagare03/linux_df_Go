package sorting

import ("strconv"
		"strings"
		"errors"
		"project/parser"
		"sort")

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

func GetTop2Size(system []parser.Fs) ([]parser.Fs, error){
	if(len(system)<=0){
		return system, errors.New("wrong input size")
	}

	sort.Slice(system, func(i, j int) bool {
		sizei := parseHumanReadableSize(system[i].Size)
		sizej := parseHumanReadableSize(system[j].Size)
		return sizei > sizej
	})
	return system, nil
}

func GetTop2Avail(system []parser.Fs) ([]parser.Fs, error){
	if(len(system)<=0){
		return system, errors.New("wrong input size")
	}
	sort.Slice(system, func(i, j int) bool {
		availi := parseHumanReadableSize(system[i].Avail)
		availj := parseHumanReadableSize(system[j].Avail)
		return availi > availj
	})
	return system, nil

}

func GetTop2Use(system []parser.Fs) ([]parser.Fs, error){
	if(len(system)<=0){
		return system, errors.New("wrong input size")
	}
	sort.Slice(system, func(i, j int) bool {
		usei, _ := strconv.ParseInt(strings.TrimSuffix(system[i].Use, "%"), 10, 64)
		usej, _ := strconv.ParseInt(strings.TrimSuffix(system[j].Use, "%"), 10, 64)
		return usei > usej
	})
	return system, nil
	
}