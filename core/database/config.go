package database

import "os"

var (
	host = os.Getenv("DB_HOST")
  port = os.Getenv("DB_PORT")
  user = os.Getenv("DB_USER")
  password = os.Getenv("DB_PASSWORD")
  name = os.Getenv("DB_NAME")
)

const (
	sslmode = "disable"
	maxRetries = 10
	delay = 5
	noEnvironmentVariablesError = "database environment variables not set"
)