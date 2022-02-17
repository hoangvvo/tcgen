package gen

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// FileIO

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("failed removing original file: %s", err)
	}
	return nil
}

func FormatPathCwd(filePath string) string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if filepath.IsAbs(filePath) {
		return filePath
	}
	return filepath.Join(cwd, filePath)
}

// String

func GetFilepath(promptStr string) string {
	var filePath string
	for len(filePath) == 0 {
		fmt.Print(promptStr)
		fmt.Scanln(&filePath)
		filePath = FormatPathCwd(filePath)
		if _, err := os.Stat(filePath); err != nil {
			LogError(err.Error())
			filePath = ""
		}
	}
	return filePath
}

func getOutFilename(basename string) string {
	return strings.TrimSuffix(filepath.Base(basename), filepath.Ext(basename))
}

func GetNumber(promptStr string) int {
	var number *int
	for number == nil {
		var inpStr string
		fmt.Print(promptStr)
		fmt.Scanln(&inpStr)
		intVal, err := strconv.Atoi(inpStr)
		if err != nil {
			LogError(err.Error())
		} else {
			number = &intVal
		}
	}
	return *number
}

// Logging

var colorError = color.New(color.FgRed)
var colorTask = color.New(color.FgMagenta)
var colorSuccess = color.New(color.FgGreen)

func LogSuccess(str string) {
	colorSuccess.Println(str)
}

func LogTask(str string) {
	colorTask.Println("[task] " + str)
}

func LogError(str string) {
	colorError.Println("[error] " + str)
}
