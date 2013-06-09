package config

func (self *Config) AddOption(section string, option string, value string) bool {
  self.AddSection(section) // Make sure section exists

	v, ok := self.data[section][option]
	if !ok {
		self.data[section][option] = &TValue{0, make([]string, 0)}
		v = self.data[section][option]
	}

	v.Count++
	v.Value = append(v.Value, value)

	return !ok
}

func (self *Config) HasOption(section string, option string) bool {
	if _, ok := self.data[section]; !ok {
		return false
	}

	_, oknd := self.data[section][option]

	return oknd
}
