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

type PokeResponse struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []interface{} `json:"results"`
}
