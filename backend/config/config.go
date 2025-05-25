package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerPort string `env:"SERVER_PORT,required"`
	DBHost     string `env:"DB_HOST,required"`
	DBName     string `env:"DB_NAME,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBSSLmode  string `env:"DB_SSLMODE,required"`
}

func NewEnConfig() *EnvConfig {
	rootPath := findRootPath()
	envPath := filepath.Join(rootPath, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("[Fatal] Unable to load the .env file from %s: %v", envPath, err)
	}

	cfg := &EnvConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("[Fatal] Unable to parse environment variables: %v", err)
	}

	return cfg
}

// This walks up the directory tree to find the project root (where .env lives)
func findRootPath() string {
	dir, _ := os.Getwd()

	for {
		if _, err := os.Stat(filepath.Join(dir, ".env")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			log.Fatal("[Fatal] Could not find .env file in any parent directory")
		}
		dir = parent
	}
}
