package config

import (
  "bufio"
	"errors"
	"io"
	"os"
	"strings"
)

func _read(fname string, c *Config) (*Config, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err = c.read(bufio.NewReader(file)); err != nil {
		return nil, err
	}

	return c, nil
}

func Read(fname string) (*Config, error) {
	return _read(fname, New())
}

var default_config *Config = New()

func ReadDefault(fname string) error {
	_, err := _read(fname, default_config)

	return err
}

func (self *Config) read(buf *bufio.Reader) (err error) {
	section := DEFAULT_SECTION
	var option string

	for {
		l, err := buf.ReadString('\n') // parse line-by-line
		if err == io.EOF {
			if len(l) == 0 {
				break
			}
		} else if err != nil {
			return err
		}

		l = strings.TrimSpace(l)

		// Switch written for readability (not performance)
		switch {
		// Empty line and comments
		case len(l) == 0, l[0] == '#', l[0] == ';':
			continue

		// New section
		case l[0] == '[' && l[len(l)-1] == ']':
			section = strings.TrimSpace(l[1 : len(l)-1])
			self.AddSection(section)

		// Other alternatives
		default:
			i := strings.IndexAny(l, "=")

			switch {
			// Option and value
			case i > 0:
				option = strings.TrimSpace(l[0:i])
				value := strings.TrimSpace((l[i+1:]))
				self.AddOption(section, option, value)
			default:
				return errors.New("could not parse line: " + l)
			}
		}
	}
	return nil
}
