package config

import "os"

// GetPort to get the port the go server will run
func GetPort() string {
	return os.Getenv("PORT")
}

// GetPort to get the port the go server will run
func GetGinMode() string {
	return os.Getenv("GIN_MODE")
}
