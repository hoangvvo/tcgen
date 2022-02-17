package gen

import (
	_ "embed"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type ConfigTestCase struct {
	Extensions []string
}

type ConfigLanguage struct {
	Name       string
	Extensions []string
	Compile    *[]string
	Run        []string
}

type ConfigOutput struct {
	Dir string
}

type Config struct {
	Testcase  ConfigTestCase
	Languages map[string]ConfigLanguage
	Output    ConfigOutput
}

//go:embed judgen.yml
var defaultConfigDat []byte

func GetConf() *Config {
	var configDat []byte

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if configDatAlt, err := os.ReadFile(path.Join(cwd, "judgen.yml")); err == nil {
		configDat = configDatAlt
	} else if os.IsNotExist(err) {
		configDat = defaultConfigDat
	} else {
		panic(err)
	}

	var conf Config
	err = yaml.Unmarshal(configDat, &conf)
	if err != nil {
		panic(err)
	}
	return &conf
}
