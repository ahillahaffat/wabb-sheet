package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    VerifyToken   string
    AccessToken   string
    PhoneNumberID string
    SpreadsheetID string
    Port          string
}

func LoadConfig() *Config {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using environment variables")
    }

    config := &Config{
        VerifyToken:   getEnv("VERIFY_TOKEN", "benefica3"),
        AccessToken:   getEnv("ACCESS_TOKEN", ""),
        PhoneNumberID: getEnv("PHONE_NUMBER_ID", ""),
        SpreadsheetID: getEnv("SPREADSHEET_ID", ""),
        Port:          getEnv("PORT", "3000"),
    }

    if config.AccessToken == "" {
        log.Fatal("ACCESS_TOKEN environment variable is required")
    }
    if config.PhoneNumberID == "" {
        log.Fatal("PHONE_NUMBER_ID environment variable is required")
    }
    if config.SpreadsheetID == "" {
        log.Fatal("SPREADSHEET_ID environment variable is required")
    }

    return config
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}