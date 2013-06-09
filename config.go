package config

import ()

const (
  DEFAULT_SECTION = "DEFAULT"
)

var (
	// Strings accepted as boolean.
	boolString = map[string]bool{
		"t":     true,
		"true":  true,
		"y":     true,
		"yes":   true,
		"on":    true,
		"1":     true,
		"f":     false,
		"false": false,
		"n":     false,
		"no":    false,
		"off":   false,
		"0":     false,
	}
)

type TValue struct {
	Count int
	Value []string
}

type Config struct {
	data map[string]map[string]*TValue
}

func New() *Config {
	c := new(Config)
	c.data = make(map[string]map[string]*TValue)
	c.AddSection(DEFAULT_SECTION) // Default section always exists.
	return c
}
