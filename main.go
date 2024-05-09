package main

import (
	"os"
	"os/exec"
	"project/logger"
	"project/parser"
	"project/sorting"
	"sync"
)

func main() {
	logger.Info("Main initialised")
	// os.Remove("dfResults.txt")

	wg := &sync.WaitGroup{}
	op := make(chan string)
	wg.Add(2)

	go func() {
		defer wg.Done()
		output, err := exec.Command("df", "-h").Output()
		op <- string(output)
		if err != nil {
			logger.Error("error executing df command :", err)
			panic(err)
		} else {
			logger.Debug("df command executed wo err")
		}
	}()

	go func() {
		defer wg.Done()
		dfop := <-op
		res := "df output \n" + dfop

		logger.Info("calling parser")
		system, err := parser.GetResults(dfop)
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
			op, _ := parser.GetFormattedDFOutput(system, 2)
			res += "\nTop 2 filesystem by size are :" + op
		}

		logger.Info("Entered sorting for Avail")

		system, err = sorting.GetTop2Avail(system)

		if err != nil {
			logger.Error("GetTop2Avail returned with error: ", err)
		} else {
			logger.Debug("GetTop2Avail returned a output")
			op, _ := parser.GetFormattedDFOutput(system, 2)
			res += "\nTop 2 filesystem by Avail are :" + op
		}

		logger.Info("Entered sorting for use")

		system, err = sorting.GetTop2Use(system)
		if err != nil {
			logger.Error("GetTop2Use returned with error: ", err)
		} else {
			logger.Debug("GetTop2Use returned a output")
			op, _ := parser.GetFormattedDFOutput(system, 2)
			res += "\nTop 2 filesystem by use are :" + op
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
	}()

	wg.Wait()

}
