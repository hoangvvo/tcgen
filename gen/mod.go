package gen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func Execute(conf *Config, genPath string, solPath string, total int) {
	rootTempDir := filepath.Join(os.TempDir(), "tcgen")
	os.MkdirAll(rootTempDir, os.ModePerm)

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	outDir := PrepareOutdir(conf)

	tempDir, err := ioutil.TempDir(rootTempDir, filepath.Base(cwd))
	// cleanup
	defer os.RemoveAll(tempDir)

	if err != nil {
		panic(err)
	}

	LogTask("Create temporary directory: " + tempDir)

	RunCmds(
		conf,
		CompileFile(conf, genPath, tempDir),
		CompileFile(conf, solPath, tempDir),
		total, tempDir, outDir)

	LogSuccess("Created " + strconv.Itoa(total) + " test cases " + outDir)
}
