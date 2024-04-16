package main

import (
	"fmt"
	"os"
	"os/exec"
	"project/logger"
	"project/parser"
	"project/sorting"
)

func main() {
	logger.Info("Main initialised")

	os.Remove("dfResults.txt")
	output, err := exec.Command("df", "-h").Output()
	op := string(output)

	if err != nil {
		logger.Error("error executing df command :", err)
		panic(err)
	} else {
		logger.Debug("df command executed wo err")
	}

	res := "df output \n" + op

	logger.Info("calling parser")

	system, err := parser.GetResults(op)
	fmt.Printf("%+v", system)
	if err != nil {
		logger.Debug("Parser returned a valid array")
	}
	logger.Info("back from parser")
	logger.Info("Entered sorting for size")

	system, err = sorting.GetTop2Size(system)

	if err != nil {
		logger.Error("GetTop2Size returned with error: ", err)
	} else {
		logger.Debug("GetTop2Size returned a output")
		res += "\nTop 2 filesystem by size are :" + parser.GetFormattedDFOutput(system, 2)
	}

	logger.Info("Entered sorting for Avail")

	system, err = sorting.GetTop2Avail(system)

	if err != nil {
		logger.Error("GetTop2Avail returned with error: ", err)
	} else {
		logger.Debug("GetTop2Avail returned a output")
		res += "\nTop 2 filesystem by Avail are :" + parser.GetFormattedDFOutput(system, 2)
	}

	logger.Info("Entered sorting for use")

	system, err = sorting.GetTop2Use(system)
	if err != nil {
		logger.Error("GetTop2Use returned with error: ", err)
	} else {
		logger.Debug("GetTop2Use returned a output")
		res += "\nTop 2 filesystem by use are :" + parser.GetFormattedDFOutput(system, 2)
	}

	file, err := os.Create("dfResults.txt")

	if err != nil {
		logger.Error("error creating output file (dfResults)", err)
		panic(err)
	} else {
		logger.Debug("output file (dfResults) created successfully")
	}

	defer file.Close()

	file.WriteString(res)

	logger.Info("output written to the file")
	logger.Info("Program executed successfully")
}
