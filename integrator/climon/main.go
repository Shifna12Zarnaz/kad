package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/kube-tarian/kad/integrator/climon/pkg/application"
	"github.com/kube-tarian/kad/integrator/climon/pkg/db/cassandra"
	"github.com/kube-tarian/kad/integrator/common-pkg/logging"
)

func main() {
	logger := logging.NewLogger()
	logger.Infof("Started deployment worker\n")

	db, err := cassandra.Create(logger)
	if err != nil {
		logger.Fatalf("failed to create db connection", err)
	}

	app := application.New(logger, db)
	go app.Start()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	app.Close()
	db.Close()
	logger.Infof("Exiting deployment worker\n")
}
