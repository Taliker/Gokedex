package commands

import "os"

func CommandExit(config *Config) error {
	os.Exit(0)
	return nil
}
