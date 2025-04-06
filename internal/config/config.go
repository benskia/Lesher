package config

import (
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

	// To work with a config, we'll need the file and the directory where it lives.
	err := os.MkdirAll(configDir, 0666)
	if err != nil {
		return nil, fmt.Errorf("NewConfig: %v", err)
	}

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

	return cfg, nil
}
