package config

import "github.com/BurntSushi/toml"

type Config struct {
	Discord       *DiscordConfig       `toml:"discord"`
	BattleMetrics *BattleMetricsConfig `toml:"battlemetrics"`
	Database      *DatabaseConfig      `toml:"database"`
}

type DiscordConfig struct {
	Token string `toml:"token"`
}

type BattleMetricsConfig struct {
	Token string `toml:"token"`
}

type DatabaseConfig struct {
	Path string `toml:"path"`
}

// Load configuration
func Load(path string) (*Config, error) {
	config := new(Config)
	_, err := toml.DecodeFile(path, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
