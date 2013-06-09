package config

type sectionError string

func (self sectionError) Error() string {
  return "section not found: " + string(self)
}

type optionError string

func (self optionError) Error() string {
	return "option not found: " + string(self)
}
