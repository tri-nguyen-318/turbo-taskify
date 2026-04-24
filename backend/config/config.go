package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port             string
	MongoURI         string
	MongoDB          string
	JWTAccessSecret  string
	JWTRefreshSecret string
	AccessTokenTTL   time.Duration
	RefreshTokenTTL  time.Duration
	GoogleClientID   string
	FrontendURL      string
}

func Load() *Config {
	accessMin, _ := strconv.Atoi(env("ACCESS_TOKEN_TTL_MINUTES", "15"))
	refreshDays, _ := strconv.Atoi(env("REFRESH_TOKEN_TTL_DAYS", "7"))

	return &Config{
		Port:             env("PORT", "8080"),
		MongoURI:         env("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:          env("MONGO_DB", "turbo_taskify"),
		JWTAccessSecret:  env("JWT_ACCESS_SECRET", "access-secret-change-me"),
		JWTRefreshSecret: env("JWT_REFRESH_SECRET", "refresh-secret-change-me"),
		AccessTokenTTL:   time.Duration(accessMin) * time.Minute,
		RefreshTokenTTL:  time.Duration(refreshDays) * 24 * time.Hour,
		GoogleClientID:   env("GOOGLE_CLIENT_ID", ""),
		FrontendURL:      env("FRONTEND_URL", "http://localhost:3000"),
	}
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
