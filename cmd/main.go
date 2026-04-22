package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jupitters/go-pizza-tracker/internal/models"
)

func main() {
	cfg := loadConfig()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dbModel, err := models.InitDB(cfg.DBPath)
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}

	slog.Info("Database initialized successfully!")

	RegisterCustomValidators()

	h := NewHandler(dbModel)

	router := gin.Default()

	if err := loadTemplates(router); err != nil {
		slog.Error("Failed to load templates", "error", err)
		os.Exit(1)
	}

	sessionStore := setupSessionStore(dbModel.Order.DB)

	setupRoutes(router, h)

	slog.Info("Server starting", "url", "http://localhost:"+cfg.Port)

	router.Run(":" + cfg.Port)
}
