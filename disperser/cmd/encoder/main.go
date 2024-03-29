package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/0glabs/0g-data-avail/common/logging"
	"github.com/0glabs/0g-data-avail/disperser/cmd/encoder/flags"
	"github.com/urfave/cli"
)

var (
	// Version is the version of the binary.
	Version   string
	GitCommit string
	GitDate   string
)

func main() {

	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "encoder"
	app.Usage = "ZGDA Encoder"
	app.Description = "Service for encoding blobs"

	app.Action = RunEncoderServer
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("application failed: %v", err)
	}

	select {}
}

func RunEncoderServer(ctx *cli.Context) error {

	config := NewConfig(ctx)

	logger, err := logging.GetLogger(config.LoggerConfig)
	if err != nil {
		return err
	}

	enc, err := NewEncoderGRPCServer(config, logger)
	if err != nil {
		return err
	}
	defer enc.Close()

	err = enc.Start(context.Background())
	if err != nil {
		return err
	}

	return nil
}
