package config

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/franciscolkdo/family-feud/internal/game"
)

//go:embed config.json
var f embed.FS

const configPath = "config.json"

var ErrorConfig = errors.New("config error")

func wrapErr(err error) error {
	return fmt.Errorf("%w: %w", ErrorConfig, err)
}

func GetConfig(path string) (game.Config, error) {
	var cfg game.Config
	var data []byte
	var err error

	if path != "" {
		data, err = os.ReadFile(path)
	} else {
		data, err = f.ReadFile(configPath)
	}

	if err != nil {
		return game.Config{}, wrapErr(err)
	}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return game.Config{}, wrapErr(err)
	}
	return cfg, nil
}
