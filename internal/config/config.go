package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

// Stores information about existing profiles and filepath(s).

type Profile struct {
	Name  string
	Start int
	End   int
}

type config struct {
	configPath string
	Profiles   []Profile
}

var Defaults []Profile = []Profile{
	{
		Name:  "mid",
		Start: 40,
		End:   50,
	},
	{
		Name:  "high",
		Start: 70,
		End:   80,
	},
}

func NewConfig() (*config, error) {
	xdgCfg := os.Getenv("XDG_CONFIG_HOME")
	configPath := path.Join(xdgCfg, "lesher/config.json")

	cfg := &config{configPath: configPath}
	if err := cfg.readConfig(); err != nil {
		cfg.Profiles = Defaults
	}

	// To work with a config, we'll need the file and directory where it lives.
	if err := cfg.writeConfig(); err != nil {
		return nil, fmt.Errorf("failed to write config file: %v", err)
	}

	return cfg, nil
}

func (cfg *config) readConfig() error {
	b, err := os.ReadFile(cfg.configPath)
	if err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}

	var profiles []Profile
	if err = json.Unmarshal(b, &profiles); err != nil {
		return fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	cfg.Profiles = profiles
	return nil
}

func (cfg *config) writeConfig() error {
	if err := os.MkdirAll(path.Dir(cfg.configPath), 0755); err != nil {
		return fmt.Errorf("error making config directories: %v", err)
	}

	f, err := os.Create(cfg.configPath)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("error creating config file: %v", err)
	}

	b, err := json.Marshal(cfg.Profiles)
	if err != nil {
		return fmt.Errorf("error marshaling json: %v", err)
	}

	if _, err := f.Write(b); err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}

	return nil
}
