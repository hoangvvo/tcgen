package gen

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func shouldCopy(conf *Config, filename string) bool {
	if len(filename) == 0 {
		return false
	}
	extWithDot := filepath.Ext(filename)
	if len(extWithDot) == 0 {
		return false
	}
	ext := extWithDot[1:]
	for _, testcaseExts := range conf.Testcase.Extensions {
		if strings.EqualFold(ext, testcaseExts) {
			return true
		}
	}
	return false
}
func CopyArtifacts(conf *Config, inDir string, outDir string) {
	files, err := ioutil.ReadDir(inDir)
	if err != nil {
		panic(err)
	}

	os.MkdirAll(outDir, os.ModePerm)

	for _, file := range files {
		filename := file.Name()
		if shouldCopy(conf, filename) {
			err := MoveFile(path.Join(inDir, filename), path.Join(outDir, filename))
			if err != nil {
				panic(err)
			}
		}
	}
}

type Executor func(extraArgs ...string) *exec.Cmd

func RunCmds(conf *Config, cmdGenExec Executor, cmdSolExec Executor, total int, inDir string, outDir string) {
	LogTask("thuc hien sinh test")
	bar := progressbar.Default(int64(total))
	for i := 0; i < total; i += 1 {
		var errGenb bytes.Buffer
		cmdGen := cmdGenExec(strconv.Itoa(i))

		cmdGen.Stderr = &errGenb
		if err := cmdGen.Run(); err != nil {
			LogError(errGenb.String())
			panic(err)
		}

		var errSolb bytes.Buffer
		cmdSol := cmdSolExec(strconv.Itoa(i))
		cmdSol.Stderr = &errSolb
		if err := cmdSol.Run(); err != nil {
			LogError(errSolb.String())
			panic(err)
		}

		CopyArtifacts(conf, inDir, path.Join(outDir, "TEST"+strconv.Itoa(i)))

		bar.Add(1)
	}
}

func PrepareOutdir(conf *Config) string {
	var outDir string
	if conf.Output.Dir == "" {
		panic(errors.New("no out.dir in config"))
	}
	outDir = FormatPathCwd(conf.Output.Dir)
	os.RemoveAll(outDir)
	os.MkdirAll(outDir, os.ModePerm)
	return outDir
}
