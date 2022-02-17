package gen

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

const SOURCEARG = "SOURCE"
const OUTPUTARG = "OUTPUT"

func getLanguage(conf *Config, ext string) *ConfigLanguage {
	for _, language := range conf.Languages {
		for _, confExt := range language.Extensions {
			if strings.EqualFold(ext, confExt) {
				return &language
			}
		}
	}
	return nil
}

func nameOrSourceReplace(arg string, path string, outPath string) string {
	if arg == SOURCEARG {
		return path
	} else if arg == OUTPUTARG {
		return outPath
	}
	return arg
}

func CompileFile(conf *Config, inPath string, outDir string) Executor {
	LogTask("Compiling: " + inPath)

	inputExt := filepath.Ext(inPath)[1:]
	lang := getLanguage(conf, inputExt)
	if lang == nil {
		panic(errors.New("\tNot supported: " + inputExt))
	} else {
		fmt.Println("\tLanguage:", lang.Name)
	}

	outFilename := getOutFilename(inPath) // get a file without ext and path
	outPath := filepath.Join(outDir, outFilename)

	if lang.Compile == nil {
		LogSuccess("\tNo compilation needed")
		outPath = outPath + "." + inputExt // add back ext
		CopyFile(inPath, outPath)
	} else {
		var cmdName string
		var cmdArgs []string
		for idx, bareCmdArg := range *lang.Compile {
			if idx == 0 {
				cmdName = nameOrSourceReplace(bareCmdArg, inPath, outPath)
			} else {
				cmdArgs = append(cmdArgs, nameOrSourceReplace(bareCmdArg, inPath, outPath))
			}
		}

		cmd := exec.Command(cmdName, cmdArgs...)

		var errb bytes.Buffer
		cmd.Stderr = &errb
		err := cmd.Run()

		if err != nil {
			LogError(errb.String())
			panic(err)
		}

		LogSuccess("\tCompilation successfully")
	}

	return func(extraArgs ...string) *exec.Cmd {
		var cmdName string
		var cmdArgs []string
		for idx, bareCmdArg := range lang.Run {
			if idx == 0 {
				cmdName = nameOrSourceReplace(bareCmdArg, inPath, outPath)
			} else {
				cmdArgs = append(cmdArgs, nameOrSourceReplace(bareCmdArg, inPath, outPath))
			}
		}
		cmdArgs = append(cmdArgs, extraArgs...)
		cmd := exec.Command(cmdName, cmdArgs...)
		cmd.Dir = outDir

		return cmd
	}
}
