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
	configDir := path.Join(os.Getenv("XDG_CONFIG_HOME"), "lesher")

	// When creating a new config, let's include some sane defaults. A
	// battery-life saving "mid" profile and a "high" profile that avoids
	// high temps at 100% charge.
	cfg := &config{
		configPath: configDir + "config.json",
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
		return nil, fmt.Errorf("createConfigFile: %v", err)
	}

	return cfg, nil
}

func (cfg *config) writeConfig() error {
	if err := os.MkdirAll(path.Dir(cfg.configPath), 0755); err != nil {
		return fmt.Errorf("writeConfig: %v", err)
	}

	f, err := os.Create(cfg.configPath)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("writeConfig: %v", err)
	}

	b, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("createConf")
	}

	return nil
}
