package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-akka/configuration"
	"github.com/go-akka/configuration/hocon"
)

var (
	emptyConfig = configuration.FromObject(hocon.NewHoconObject())
)

func LoadFromFile(path string) (cfg *configuration.Config, err error) {
	defer func() {
		if errIn := recover(); errIn != nil {
			cfg = nil
			err = fmt.Errorf("loading config '%v' failed: %w", path, errIn)
		}
	}()

	//if the path doesn't exist or it isn't a file, returns empty config
	if info, errIn := os.Stat(path); errIn != nil || info.IsDir() {
		if errIn != nil {
			err = fmt.Errorf("couldn't read config from %v: %w", path, errIn)
		} else {
			err = fmt.Errorf("couldn't read config, %v is a directory", path)
		}
		cfg = nil
		return
	}

	cfg = configuration.LoadConfig(path)
	err = nil
	return
}

func GetConfig(c *configuration.Config, path string) (cfg *configuration.Config, err error) {
	defer func() {
		if errIn := recover(); errIn != nil {
			cfg = nil
			err = fmt.Errorf("getting config failed: %w", errIn)
		}
	}()

	cfg = emptyConfig
	err = nil

	if scfg := c.GetConfig(path); scfg != nil {
		cfg = scfg
	}
	return
}

func LoadDefault() (cfg *configuration.Config, err error) {
	configFile := strings.TrimSuffix(
		filepath.Base(os.Args[0]), filepath.Ext(os.Args[0])) + ".conf"
	return LoadFromFile(configFile)
}
