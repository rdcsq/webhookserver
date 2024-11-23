package structs

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var Env WSEnv

type WSEnv struct {
	ListeningAddress string
	JwtSecret        []byte
	ConfigPath       string
}

func InitializeEnv() {
	ListeningAddress := os.Getenv("WS_LISTENING_ADDRESS")
	JwtSecret := os.Getenv("WS_JWT_SECRET")
	ConfigPath := os.Getenv("WS_CONFIG_PATH")

	if JwtSecret == "" {
		panic("WS_JWT_SECRET is required")
	}

	if ListeningAddress == "" {
		ListeningAddress = "0.0.0.0:3000"
	}

	if ConfigPath == "" {
		ConfigPath = "config.json"
	}

	Env = WSEnv{
		ListeningAddress: ListeningAddress,
		JwtSecret:        []byte(JwtSecret),
		ConfigPath:       ConfigPath,
	}
}
