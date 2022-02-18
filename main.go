package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hoangvvo/tcgen/gen"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			gen.LogError(r.(error).Error())
			os.Exit(1)
		}
	}()
	fmt.Print("tcgen: generate test cases for coding problems\nhttps://github.com/hoangvvo/tcgen\n\n")
	conf := gen.GetConf()
	genPath := gen.GetFilepath("Enter case generation file: ")
	solPath := gen.GetFilepath("Enter case solver file: ")
	total := gen.GetNumber("Number of run: ")

	startTime := time.Now()

	gen.Execute(conf, genPath, solPath, total)

	gen.LogSuccess("Took " + fmt.Sprint(time.Since(startTime)))
}
