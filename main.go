package main

import (
	"fmt"
	"github.com/tomdionysus/vincent"
	"github.com/tomdionysus/vincent-demo/config"
	"github.com/tomdionysus/vincent/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// This is for detecting SIGTERM (shutdown) from the OS.
	var sigint chan (os.Signal)

	// Load Config. See the config/ subpackage for more details
	cfg := config.NewConfig()
	errs := cfg.Validate()
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Printf("Config: %s\n", err)
		}
		return
	}

	// Boot the logger at the configured logging level
	logger := log.NewConsoleLogger(*cfg.LogLevel)

	// Startup Header for log
	logger.Raw(">>>>> vincent-demo v%s <<<<<", Version)
	logger.Raw("%s", time.Now())

	// Create vincent server
	logger.Debug("Creating Server")
	svr, err := vincent.New(logger)
	if err != nil {
		logger.Error("Cannot Create Server: %s", err)
		return
	}

	// Load templates from the templates/ directory
	logger.Debug("Loading Templates")
	err = svr.LoadTemplates("", "templates")
	if err != nil {
		logger.Error("Cannot load Templates: %s", err)
		return
	}

	// This is an example controller
	svr.AddController("/", func(context *vincent.Context) (bool, error) {
		context.Output["version"] = fmt.Sprintf("%s", Version)
		context.Output["port"] = *cfg.HTTPPort
		return true, nil
	})

	// Start the server on the configured port
	logger.Debug("Starting HTTP Server")
	svr.Start(fmt.Sprintf(":%d", *cfg.HTTPPort))
	logger.Info("HTTP Listening on port %d", *cfg.HTTPPort)

	// Link the OS Interrupt and SIGTERM to our listening channel
	logger.Debug("Notifying SIGINT")
	sigint = make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, syscall.SIGTERM)
	logger.Info("Started")

	// Wait for Interrupt / SIGINT
	for {
		select {
		case sig := <-sigint:
			switch sig {
			case os.Interrupt:
				fallthrough
			case syscall.SIGTERM:
				logger.Info("Signal %d received, shutting down", sig)
				goto stop
			}
		}
	}

	// Shutdown
stop:
	logger.Info("Shutdown")
}
