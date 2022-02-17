package gen

import (
	_ "embed"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ConfigTestCase struct {
	Extensions []string
	Output     string
}

type ConfigLanguage struct {
	Name       string
	Extensions []string
	Compile    *[]string
	Run        []string
}

type Config struct {
	Testcase  ConfigTestCase
	Languages map[string]ConfigLanguage
}

//go:embed judgen.yml
var defaultConfigDat []byte

func GetConf() *Config {
	var configDat []byte

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if configDatAlt, err := os.ReadFile(filepath.Join(cwd, "judgen.yml")); err == nil {
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
