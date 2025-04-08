package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/DragonPow/movie_booking/internal/auth/config"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Start the authentication service",
		RunE:  runServer,
	}

	cmd.PersistentFlags().String("config", "config.yaml", "path to config file")

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func runServer(cmd *cobra.Command, args []string) error {
	configFile, _ := cmd.Flags().GetString("config")

	// Load configuration
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		return err
	}

	// Create and start server
	srv, err := newAuthServer(cfg)
	if err != nil {
		return err
	}
	defer srv.stop()

	// Handle graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh

		log.Println("Shutting down gracefully...")
		srv.stop()
		log.Println("Server stopped")
		os.Exit(0)
	}()

	return srv.start()
}
