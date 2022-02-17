package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"

	"github.com/hoangvvo/judgen/gen"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			gen.LogError(r.(error).Error())
			os.Exit(1)
		}
	}()
	fmt.Print("Chuong trinh tao test case don gian\nhttps://github.com/hoangvvo/judgen\n\n")
	conf := gen.GetConf()

	genPath := gen.GetFilepath("Nhap file sinh test: ")
	solPath := gen.GetFilepath("Nhap file bai giai: ")
	total := gen.GetNumber("Nhap so lan chay: ")

	rootTempDir := path.Join(os.TempDir(), "judgen")
	os.MkdirAll(rootTempDir, os.ModePerm)

	wd, _ := os.Getwd()

	outDir := gen.PrepareOutdir(conf)

	tempDir, err := ioutil.TempDir(rootTempDir, path.Base(wd))
	// cleanup
	defer os.RemoveAll(tempDir)

	if err != nil {
		panic(err)
	}

	gen.LogTask("tao thu muc tam: " + tempDir)

	gen.RunCmds(
		conf,
		gen.CompileFile(conf, genPath, tempDir),
		gen.CompileFile(conf, solPath, tempDir),
		total, tempDir, outDir)

	gen.LogSuccess("tao thanh cong " + strconv.Itoa(total) + " test case trong " + outDir)
}
