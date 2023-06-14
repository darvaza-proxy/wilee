// Package main provides the Wilee server
package main

import (
	"fmt"
	"os"

	"darvaza.org/slog"
	"github.com/spf13/cobra"

	"darvaza.org/wilee/cmd/wilee/zerolog"
)

const (
	// CmdName specifies the name of the command
	CmdName = "wilee"
	// DefaultConfigFile specifies the default filename of the config file
	DefaultConfigFile = CmdName + ".toml"
	// DefaultLogLevel specifies the log level we handle by default
	DefaultLogLevel = slog.Debug
)

var (
	cfg     Config
	cfgFile string
	log     slog.Logger
)

var rootCmd = &cobra.Command{
	Use:   CmdName,
	Short: CmdName + ", ACME Server/Proxy powered by darvaza.org",
}

func fatal(err error, msg string, args ...any) {
	l := log.Fatal()

	if err != nil {
		l = l.WithField(slog.ErrorFieldName, err)
	}

	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	l.Print(msg)

	panic("unreachable")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fatal(err, "")
	}
}

// cobraInit loads the config file before calling the
// commands
func cobraInit() {
	if cfgFile != "" {
		var c Config

		err := c.ReadInFile(cfgFile)
		switch {
		case err == nil:
			// good config
			cfg = c
			return
		case os.IsNotExist(err) && cfgFile == DefaultConfigFile:
			// missing DefaultConfigFile, ignore
		default:
			// any other error is fatal
			fatal(err, "failed processing %q", cfgFile)
		}
	}

	// didn't load, apply defaults
	if err := cfg.Prepare(); err != nil {
		fatal(err, "failed to set config defaults")
	}
}

// init initialises the global logger and config-file loading
func init() {
	log = zerolog.NewLogger(DefaultLogLevel)

	// root level flags
	pflags := rootCmd.PersistentFlags()
	pflags.StringVarP(&cfgFile, "config-file", "f", DefaultConfigFile, "config file (TOML format)")

	// load config-file before the rest of the cobra commands
	cobra.OnInitialize(cobraInit)
}
