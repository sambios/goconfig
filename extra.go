package goconfig

// newConfigFile creates an empty configuration representation.
func EmptyConfigFile() *ConfigFile {
	c := new(ConfigFile)
	c.data = make(map[string]map[string]string)
	c.keyList = make(map[string][]string)
	c.sectionComments = make(map[string]string)
	c.keyComments = make(map[string]map[string]string)
	c.BlockMode = true
	return c
}