//author gonghh
//Copyright 2013 gonghh(screscent). Under the Apache License, Version 2.0.
package config

func (self *Config) AddSection(section string) bool {
  if section == "" {
		return false
	}

	if _, ok := self.data[section]; ok {
		return false
	}

	self.data[section] = make(map[string]*TValue)
	return true
}

func (self *Config) HasSection(section string) bool {
	_, ok := self.data[section]

	return ok
}
