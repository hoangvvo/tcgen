package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/hoangvvo/judgen/gen"
)

func execute(conf *gen.Config, genPath string, solPath string, total int) {
	rootTempDir := path.Join(os.TempDir(), "judgen")
	os.MkdirAll(rootTempDir, os.ModePerm)

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	outDir := gen.PrepareOutdir(conf)

	tempDir, err := ioutil.TempDir(rootTempDir, path.Base(cwd))
	// cleanup
	defer os.RemoveAll(tempDir)

	if err != nil {
		panic(err)
	}

	gen.LogTask("Create temporary directory: " + tempDir)

	gen.RunCmds(
		conf,
		gen.CompileFile(conf, genPath, tempDir),
		gen.CompileFile(conf, solPath, tempDir),
		total, tempDir, outDir)

	gen.LogSuccess("Created " + strconv.Itoa(total) + " test cases " + outDir)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			gen.LogError(r.(error).Error())
			os.Exit(1)
		}
	}()
	fmt.Print("judgen: generate test cases for coding problems\nhttps://github.com/hoangvvo/judgen\n\n")
	conf := gen.GetConf()
	genPath := gen.GetFilepath("Enter case generation file: ")
	solPath := gen.GetFilepath("Enter case solver file: ")
	total := gen.GetNumber("Number of run: ")

	startTime := time.Now()

	execute(conf, genPath, solPath, total)

	gen.LogSuccess("Took " + fmt.Sprint(time.Since(startTime)))
}
