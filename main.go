package main

import (
	"os"
	"os/exec"
	"project/logger"
	"project/parser"
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
	res += parser.GetResults(op)
	logger.Info("back from parser")
	
	if(len(res)!=0){
		logger.Debug("Parser returned a non zero length string")
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
