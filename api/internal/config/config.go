package config

import "net/http"

type Config struct {
	CORSorigins []string
	CORSmethods []string
	DBName      string
	Host        string
	Port        string
}

func Load() *Config {
	c := Config{
		CORSorigins: []string{
			"http://localhost:8080",
			"http://localhost:3000",
		},
		CORSmethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		DBName: "tamanotchi.sqlite",
		Host:   "0.0.0.0",
		Port:   "8080",
	}
	return &c
}
