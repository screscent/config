package config

import (
  "errors"
	"strconv"
	"strings"
)

func (self *Config) Bool(section string, option string) (value bool, err error) {
	sv, err := self.String(section, option)
	if err != nil {
		return false, err
	}

	value, ok := boolString[strings.ToLower(sv)]
	if !ok {
		return false, errors.New("could not parse bool value: " + sv)
	}

	return value, nil
}

func (self *Config) BoolDefault(option string) (value bool, err error) {
	return self.Bool(DEFAULT_SECTION, option)
}

func Bool(section string, option string) (value bool, err error) {
	return default_config.Bool(section, option)
}

func BoolDefault(option string) (value bool, err error) {
	return default_config.Bool(DEFAULT_SECTION, option)
}

func (self *Config) Float(section string, option string) (value float64, err error) {
	sv, err := self.String(section, option)
	if err == nil {
		value, err = strconv.ParseFloat(sv, 64)
	}

	return value, err
}

func (self *Config) FloatDefault(option string) (value float64, err error) {
	return self.Float(DEFAULT_SECTION, option)
}

func Float(section string, option string) (value float64, err error) {
	return default_config.Float(section, option)
}

func FloatDefault(option string) (value float64, err error) {
	return default_config.Float(DEFAULT_SECTION, option)
}

func (self *Config) Int(section string, option string) (value int, err error) {
	sv, err := self.String(section, option)
	if err == nil {
		value, err = strconv.Atoi(sv)
	}

	return value, err
}

func (self *Config) IntDefault(option string) (value int, err error) {
	return self.Int(DEFAULT_SECTION, option)
}

func Int(section string, option string) (value int, err error) {
	return default_config.Int(section, option)
}
func IntDefault(option string) (value int, err error) {
	return default_config.Int(DEFAULT_SECTION, option)
}

func (self *Config) String(section string, option string) (value string, err error) {
	if _, ok := self.data[section]; ok {
		if v, ok := self.data[section][option]; ok {
			if v.Count > 0 {
				return v.Value[0], nil
			}
			return "", errors.New(optionError(section + "::" + option).Error())
		}

		return "", errors.New(optionError(section + "::" + option).Error())
	}

	return "", errors.New(sectionError(section).Error())
}

func (self *Config) StringDefault(option string) (value string, err error) {
	return self.String(DEFAULT_SECTION, option)
}

func String(section string, option string) (value string, err error) {
	return default_config.String(section, option)
}

func StringDefault(option string) (value string, err error) {
	return default_config.String(DEFAULT_SECTION, option)
}

func (self *Config) Strings(section string, option string) (value []string, err error) {
	if _, ok := self.data[section]; ok {
		if v, ok := self.data[section][option]; ok {
			if v.Count > 0 {
				return v.Value, nil
			}
			return nil, errors.New(optionError(section + "::" + option).Error())
		}

		return nil, errors.New(optionError(section + "::" + option).Error())
	}

	return nil, errors.New(sectionError(section).Error())
}

func (self *Config) StringsDefault(option string) (value []string, err error) {
	return self.Strings(DEFAULT_SECTION, option)
}

func Strings(section string, option string) (value []string, err error) {
	return default_config.Strings(section, option)
}

func StringsDefault(option string) (value []string, err error) {
	return default_config.Strings(DEFAULT_SECTION, option)
}
