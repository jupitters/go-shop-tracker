package main

import (
	"encoding/json"
	"html/template"
	"os"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Config struct {
	Port   string
	DBPath string
}

func loadConfig() Config {
	return Config{
		Port:   getEnv("PORT", "8080"),
		DBPath: getEnv("DATABASE_URL", "./data/orders.db"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func loadTemplates(router *gin.Engine) error {
	functions := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"json": func(v interface{}) template.JS {
			b, _ := json.Marshal(v)
			return template.JS(b)
		},
	}

	tmpl, err := template.New("").Funcs(functions).ParseGlob("templates/*.tmpl")
	if err != nil {
		return err
	}
	router.SetHTMLTemplate(tmpl)
	return nil
}

func setupSessionStore(db *gorm.DB, secretKey []byte) sessions.Store {
	store := gormsessions.NewStore(db, true, secretKey)
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
		Secure:   true,
		SameSite: 3,
	})

	return store
}
