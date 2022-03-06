package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/hoangvvo/tcgen/gen"
	"gopkg.in/yaml.v2"
)

//go:embed tcgen.yml
var defaultConfigDat []byte

func getConf() *gen.Config {
	var configDat []byte

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if configDatAlt, err := os.ReadFile(filepath.Join(cwd, "tcgen.yml")); err == nil {
		configDat = configDatAlt
	} else if os.IsNotExist(err) {
		configDat = defaultConfigDat
	} else {
		panic(err)
	}

	var conf gen.Config
	err = yaml.Unmarshal(configDat, &conf)
	if err != nil {
		panic(err)
	}
	return &conf
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			gen.LogError(r.(error).Error())
			os.Exit(1)
		}
	}()
	fmt.Print("tcgen: generate test cases for coding problems\nhttps://github.com/hoangvvo/tcgen\n\n")
	conf := getConf()
	genPath := gen.GetFilepath("Enter case generation file: ")
	solPath := gen.GetFilepath("Enter case solver file: ")
	total := gen.GetNumber("Number of run: ")

	startTime := time.Now()

	gen.Execute(conf, genPath, solPath, total)

	gen.LogSuccess("Took " + fmt.Sprint(time.Since(startTime)))

	fmt.Print("Press 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
