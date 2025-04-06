package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

// Stores information about existing profiles and filepath(s).

type config struct {
	configPath string
	Profiles   []Profile
}

type Profile struct {
	Name  string
	Start int
	End   int
}

func NewConfig() (*config, error) {
	xdgCfg := os.Getenv("XDG_CONFIG_HOME")
	configPath := path.Join(xdgCfg, "lesher/config.json")

	// When creating a new config, let's include some sane defaults. A
	// battery-life saving "mid" profile and a "high" profile that avoids
	// high temps at 100% charge.
	cfg := &config{
		configPath: configPath,
		Profiles: []Profile{
			{
				Name:  "mid",
				Start: 40,
				End:   50,
			},
			{
				Name:  "high",
				Start: 80,
				End:   90,
			},
		},
	}

	// To work with a config, we'll need the file and directory where it lives.
	if err := cfg.writeConfig(); err != nil {
		return nil, fmt.Errorf("failed to write config file: %v", err)
	}

	return cfg, nil
}

func (cfg *config) writeConfig() error {
	if err := os.MkdirAll(path.Dir(cfg.configPath), 0755); err != nil {
		return err
	}

	f, err := os.Create(cfg.configPath)
	defer f.Close()
	if err != nil {
		return err
	}

	b, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	if _, err := f.Write(b); err != nil {
		return err
	}

	return nil
}
