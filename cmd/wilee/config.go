package main

import (
	"bytes"
	"io"

	"github.com/BurntSushi/toml"

	"darvaza.org/darvaza/shared/config"
	"darvaza.org/darvaza/shared/config/expand"
)

// Config represents the configuration file for setting up
// the wilee Server
type Config struct{}

// Prepare attempts to validate and fill the gaps on a
// Config object
func (c *Config) Prepare() error {
	return config.Prepare(c)
}

// ReadInFile loads a TOML config file into a Config object,
// validates and fills the gaps
func (c *Config) ReadInFile(filename string) error {
	// read and expand
	s, err := expand.FromFile(filename, nil)
	if err != nil {
		return err
	}

	// decode
	_, err = toml.Decode(s, c)
	if err != nil {
		return err
	}

	// and validate
	return c.Prepare()
}

// WriteTo writes the Config in TOML format
func (c *Config) WriteTo(w io.Writer) (int64, error) {
	var buf bytes.Buffer

	// encode
	enc := toml.NewEncoder(&buf)
	err := enc.Encode(c)
	if err != nil {
		return 0, err
	}

	// write
	return buf.WriteTo(w)
}
