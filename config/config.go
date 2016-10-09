package config

import (
	"errors"
	"flag"
	"os"
	"strconv"
)

type Config struct {
	LogLevel *string
	HTTPPort *uint
}

func NewConfig() *Config {
	c := &Config{}

	// Get the PORT from the env, default to 8080. Heroku and other PaaS providers set PORT.
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	c.LogLevel = flag.String("loglevel", os.Getenv("LOG_LEVEL"), "Logging Level [error,warn,info,debug]")
	c.HTTPPort = flag.Uint("port", uint(port), "HTTP port")
	flag.Parse()

	return c
}

func (me *Config) Validate() []error {
	out := []error{}

	if *me.LogLevel == "" {
		out = append(out, errors.New("-loglevel or ENV LOG_LEVEL must be specified"))
	}
	if me.HTTPPort == nil {
		out = append(out, errors.New("-port or ENV PORT must be specified"))
	}

	return out
}
