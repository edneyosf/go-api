package database

var (
	host string
  port string
  user string
  password string
  name string
)

const (
	sslmode = "disable"
	maxRetries = 10
	delay = 5
	noEnvironmentVariablesError = "database environment variables not set"
)