package configstore

import (
	"errors"
	"os"
	"path/filepath"

	nanoclawconfig "github.com/probnotas/nanoClaw/pkg/config"
)

const (
	configDirName  = ".nanoclaw"
	configFileName = "config.json"
)

func ConfigPath() (string, error) {
	dir, err := ConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, configFileName), nil
}

func ConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configDirName), nil
}

func Load() (*nanoclawconfig.Config, error) {
	path, err := ConfigPath()
	if err != nil {
		return nil, err
	}
	return nanoclawconfig.LoadConfig(path)
}

func Save(cfg *nanoclawconfig.Config) error {
	if cfg == nil {
		return errors.New("config is nil")
	}
	path, err := ConfigPath()
	if err != nil {
		return err
	}
	return nanoclawconfig.SaveConfig(path, cfg)
}
