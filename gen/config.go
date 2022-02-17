package gen

import (
	"os"

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

func GetConf() *Config {
	configDat, err := os.ReadFile("judgen.yml")
	if err != nil {
		panic(err)
	}
	var conf Config
	err = yaml.Unmarshal(configDat, &conf)
	if err != nil {
		panic(err)
	}
	return &conf
}
