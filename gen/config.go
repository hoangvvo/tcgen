package gen

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
