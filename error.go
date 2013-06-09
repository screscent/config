//modify by author gonghh
//Copyright 2013 gonghh(screscent). Under the Apache License, Version 2.0.
package config

type sectionError string

func (self sectionError) Error() string {
  return "section not found: " + string(self)
}

type optionError string

func (self optionError) Error() string {
	return "option not found: " + string(self)
}
