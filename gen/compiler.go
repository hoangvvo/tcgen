package gen

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

const SOURCE_TOKEN = "$SOURCE"
const OUTPUT_TOKEN = "$OUTPUT"

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

func replaceSO(arg string, inPath string, outPath string) string {
	arg = strings.Replace(arg, SOURCE_TOKEN, inPath, -1)
	arg = strings.Replace(arg, OUTPUT_TOKEN, outPath, -1)
	return arg
}

func CompileFile(conf *Config, inPath string, outDir string) Executor {
	LogTask("Compiling: " + inPath)

	inputExt := filepath.Ext(inPath)
	if len(inputExt) > 0 {
		inputExt = inputExt[1:] // remove "."
	}
	lang := getLanguage(conf, inputExt)
	if lang == nil {
		panic(errors.New("\tNot supported: " + inputExt))
	} else {
		fmt.Println("\tLanguage:", lang.Name)
	}

	outFilename := getFilenameWithoutExt(inPath) // get a file without ext and path
	outPath := filepath.Join(outDir, outFilename)

	if lang.Compile == nil {
		LogSuccess("\tNo compilation needed")
		if len(inputExt) > 0 {
			outPath = outPath + "." + inputExt // add back ext
		}
		CopyFile(inPath, outPath)
	} else {
		var cmdName string
		var cmdArgs []string
		for idx, bareCmdArg := range *lang.Compile {
			if idx == 0 {
				cmdName = replaceSO(bareCmdArg, inPath, outPath)
			} else {
				cmdArgs = append(cmdArgs, replaceSO(bareCmdArg, inPath, outPath))
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
				cmdName = replaceSO(bareCmdArg, inPath, outPath)
			} else {
				cmdArgs = append(cmdArgs, replaceSO(bareCmdArg, inPath, outPath))
			}
		}
		cmdArgs = append(cmdArgs, extraArgs...)
		cmd := exec.Command(cmdName, cmdArgs...)
		cmd.Dir = outDir

		return cmd
	}
}
