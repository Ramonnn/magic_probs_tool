#!/bin/bash

# Replace "your-project-name" with your project name
PROJECT_NAME="go_magic_probs_tool"

# Create main project directory
# mkdir -p $PROJECT_NAME && cd $PROJECT_NAME

# Initialize Go module
go mod init $PROJECT_NAME

# Create cmd directory and main app entry
mkdir -p cmd/web
cat > cmd/web/main.go <<EOF
package main

import "fmt"

func main() {
    fmt.Println("Web app starting...")
}
EOF

# Create config directory
mkdir config
cat > config/config.go <<EOF
package config

import (
    "github.com/spf13/viper"
    "log"
)

func LoadConfig() {
    viper.SetConfigName(".env")
    viper.SetConfigType("env")
    viper.AddConfigPath(".")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error loading config file: %v", err)
    }
}
EOF

# Create internal directory structure
mkdir -p internal/{app/{calculate,fetch,print},database,models,services,handlers}

# Create database handler (PostgreSQL)
cat > internal/database/postgres.go <<EOF
package database

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
    "log"
)

var DB *pgxpool.Pool

func InitDatabase(dsn string) {
    var err error
    DB, err = pgxpool.New(context.Background(), dsn)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
    }
}
EOF

# Create models (CardData)
cat > internal/models/card.go <<EOF
package models

type CardData struct {
    UUID        string
    Name        string
    Number      string
    FrameEffects *string
    PromoTypes  *string
}
EOF

# Create a basic service file for card calculations
cat > internal/services/card_service.go <<EOF
package services

import (
    "context"
    "your-project-name/internal/database"
    "your-project-name/internal/models"
)

func FetchCardData(ctx context.Context, cardNames []string) ([]models.CardData, error) {
    query := "SELECT uuid, name, number, frameeffects, promotypes FROM cards WHERE name = ANY($1);"
    rows, err := database.DB.Query(ctx, query, cardNames)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var cards []models.CardData
    for rows.Next() {
        var card models.CardData
        if err := rows.Scan(&card.UUID, &card.Name, &card.Number, &card.FrameEffects, &card.PromoTypes); err != nil {
            return nil, err
        }
        cards = append(cards, card)
    }

    return cards, nil
}
EOF

# Create HTTP handlers
cat > internal/handlers/card_handler.go <<EOF
package handlers

import (
    "encoding/json"
    "net/http"
    "your-project-name/internal/services"
)

func GetCardDataHandler(w http.ResponseWriter, r *http.Request) {
    cardNames := []string{"Card1", "Card2"} // Example, replace with query params
    cards, err := services.FetchCardData(r.Context(), cardNames)
    if err != nil {
        http.Error(w, "Failed to fetch cards", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cards)
}
EOF

# Create pkg directory for utilities
mkdir -p pkg/utils
cat > pkg/utils/error_handler.go <<EOF
package utils

import (
    "log"
    "net/http"
)

func ErrorResponse(w http.ResponseWriter, message string, statusCode int) {
    w.WriteHeader(statusCode)
    w.Write([]byte(message))
    log.Printf("Error %d: %s", statusCode, message)
}
EOF

# Create web directory (for future frontend, static files)
mkdir -p web/{static,templates}

# Create Docker setup
cat > Dockerfile <<EOF
# Start from the official Go image
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o web cmd/web/main.go

EXPOSE 8080
CMD ["./web"]
EOF

cat > docker-compose.yml <<EOF
version: "3.8"

services:
  web:
    build: .
    ports:
      - "8080:8080"
    environment:
      - POSTGRES_USER=yourusername
      - POSTGRES_PASSWORD=yourpassword
      - POSTGRES_DB=yourdb
      - POSTGRES_HOST=yourdbhost
      - POSTGRES_PORT=5432
EOF

# Initialize Git
git init
echo -e "web/\n*.env\ndist/\n" > .gitignore

# Output success message
echo "✅ Project structure set up at $PROJECT_NAME"

