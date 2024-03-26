package commands

import "os"

func CommandExit(config *Config, arg string) error {
	os.Exit(0)
	return nil
}
