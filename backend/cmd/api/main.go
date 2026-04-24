package main

import (
	"log"

	"github.com/joho/godotenv"

	authapp "turbo-taskify/backend/internal/application/auth"
	"turbo-taskify/backend/config"
	mongoinfra "turbo-taskify/backend/internal/infrastructure/mongo"
	tokeninfra "turbo-taskify/backend/internal/infrastructure/token"
	httpserver "turbo-taskify/backend/internal/interfaces/http"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	db, err := mongoinfra.Connect(cfg.MongoURI, cfg.MongoDB)
	if err != nil {
		log.Fatalf("mongodb: %v", err)
	}

	userRepo := mongoinfra.NewUserRepository(db)
	tokenSvc := tokeninfra.NewService(cfg.JWTAccessSecret, cfg.JWTRefreshSecret, cfg.AccessTokenTTL, cfg.RefreshTokenTTL)
	authSvc := authapp.NewService(userRepo, tokenSvc, cfg.GoogleClientID)

	r := httpserver.NewServer(cfg.FrontendURL, authSvc, tokenSvc)

	log.Printf("🚀  API listening on %s", httpserver.ListenAddr(cfg.Port))
	if err := r.Run(httpserver.ListenAddr(cfg.Port)); err != nil {
		log.Fatalf("server: %v", err)
	}
}
