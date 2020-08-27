package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"go.sirus.dev/nats-logger/pkg/connector"
	"go.sirus.dev/nats-logger/pkg/logger"
	"go.sirus.dev/nats-logger/pkg/utils"
)

var rootCmd = &cobra.Command{
	Use:   "nats-logger",
	Short: "service to log nats event on database",
	Run: func(cmd *cobra.Command, args []string) {
		// initiate config
		var err error
		conf, err := LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configurations %v", err)
		}

		// setup logger
		lg, err := utils.CreateLogger(conf.LogLevel)
		if err != nil {
			log.Fatalf("Error setup logger %v", err)
		}
		lg.Info("logger setup finish")

		// connect to postgres
		models := []interface{}{}
		models = append(models, logger.Models...)
		db, err := connector.ConnectToPostgres(conf.Postgres, models)
		if err != nil {
			lg.Fatalf("failed to open postgres -> %v", err)
		}
		lg.Info("postgres connected")

		// connect to nats
		lg.Debug("connect to nats")
		nc, err := nats.Connect(conf.NatsURL)
		if err != nil {
			lg.Fatalf("failed to connect to nats", err)
		}
		lg.Debug("encode nats connecting")
		natsConn, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
		if err != nil {
			lg.Fatalf("failed to encode nats connection", err)
		}
		lg.Info("nats connected")

		// create logger instance
		lg.Debug("create logger instance")
		svc := logger.New(db, natsConn, lg)

		lg.Info("run logger service")
		err = svc.Run()
		if err != nil {
			lg.Fatalf("failed to run logger service", err)
		}
	},
}

func init() {
	rootCmd.Version = Version
}

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
