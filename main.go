package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"diet_diary/internal/config"
	"diet_diary/internal/database"
	"diet_diary/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = database.InitDB(cfg.DB)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer database.CloseDB()

	// Webservice
	// For the start, a single user is enough
	router := gin.Default()
	handlers.SetupRoutes(router)

	srv := &http.Server{
		Addr:    cfg.Srv.Host(),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds")

	log.Println("Server exiting")
}
