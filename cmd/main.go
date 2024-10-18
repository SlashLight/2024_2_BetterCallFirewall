package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx"
	"github.com/joho/godotenv"

	"github.com/2024_2_BetterCallFirewall/internal/auth/controller"
	"github.com/2024_2_BetterCallFirewall/internal/auth/repository/postgres"
	"github.com/2024_2_BetterCallFirewall/internal/auth/service"
	postController "github.com/2024_2_BetterCallFirewall/internal/post/controller"
	"github.com/2024_2_BetterCallFirewall/internal/post/repository"
	postServ "github.com/2024_2_BetterCallFirewall/internal/post/service"
	profileController "github.com/2024_2_BetterCallFirewall/internal/profile/controller"
	profileRepository "github.com/2024_2_BetterCallFirewall/internal/profile/repository"
	profileService "github.com/2024_2_BetterCallFirewall/internal/profile/service"
	"github.com/2024_2_BetterCallFirewall/internal/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	postgresDB, err := postgres.StartPostgres(connStr)
	if err != nil {
		log.Fatalf("Error starting postgres: %v", err)
	}

	repo := postgres.NewAdapter(postgresDB)
	profileRepo := profileRepository.NewProfileRepo(postgresDB)
	err = repo.CreateNewSessionTable()
	if err != nil {
		log.Fatalf("Error creating session table: %v", err)
	}
	authServ := service.NewAuthServiceImpl(repo)
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	responder := router.NewResponder(logger)
	sessionManager := service.NewSessionManager(repo)
	control := controller.NewAuthController(responder, authServ, sessionManager)

	postRepo := repository.NewRepository()
	postService := postServ.NewPostServiceImpl(postRepo)
	postControl := postController.NewPostController(postService, responder)

	profileUsecase := profileService.NewProfileUsecase(profileRepo, postRepo)
	profileControl := profileController.NewProfileController(profileUsecase, responder)

	rout := router.NewRouter(control, profileControl, postControl, sessionManager)
	server := http.Server{
		Addr:         ":8080",
		Handler:      rout,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("Starting server on port 8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
