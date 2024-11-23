package structs

import (
	"encoding/json"
	"os"
)

var Config []WSConfig

type WSConfig struct {
	Name             string   `json:"name"`
	Command          string   `json:"command"`
	Args             []string `json:"args"`
	Environment      []string `json:"environment"`
	WorkingDirectory string   `json:"workingDirectory"`
	Timeout          int      `json:"timeout"`
}

func ParseConfig() bool {
	configString, err := os.ReadFile(Env.ConfigPath)
	if err != nil {
		if Config == nil {
			panic("Config file could not be read")
		}
		return false
	}

	err = json.Unmarshal(configString, &Config)
	return err == nil
}
