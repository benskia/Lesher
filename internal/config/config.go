package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

// Description:
//	Manages information about config files and existing profiles.
//
// Responsibilities:
//	- Create config file if it doesn't exist.
//	- Read profiles from existing config file.
//	- Write profiles to config file.
//	- Store Profiles during program execution.

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

// Get a config pointer whether a previous config exists or not.
func LoadConfig() (*config, error) {
	xdgCfg := os.Getenv("XDG_CONFIG_HOME")
	configPath := path.Join(xdgCfg, "lesher/config.json")
	cfg := &config{configPath: configPath}

	// We can still use Defaults if we fail to get Profiles from a config file.
	// Still, if we were expecting success, so return that error later.
	err := cfg.readConfig()
	if err != nil {
		cfg.Profiles = Defaults
		err = fmt.Errorf("failed to read config file: %v", err)
	}

	return cfg, err
}

// Write config's profiles to file.
func (cfg *config) SaveConfig() error {
	// To work with a config, we'll need the file and directory where it lives.
	if err := cfg.writeConfig(); err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}
	return nil
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
