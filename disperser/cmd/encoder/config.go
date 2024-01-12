package main

import (
	"github.com/urfave/cli"
	"github.com/zero-gravity-labs/zerog-data-avail/common/logging"
	"github.com/zero-gravity-labs/zerog-data-avail/core/encoding"
	"github.com/zero-gravity-labs/zerog-data-avail/disperser/cmd/encoder/flags"
	"github.com/zero-gravity-labs/zerog-data-avail/disperser/encoder"
)

type Config struct {
	EncoderConfig encoding.EncoderConfig
	LoggerConfig  logging.Config
	ServerConfig  *encoder.ServerConfig
	MetricsConfig encoder.MetrisConfig
}

func NewConfig(ctx *cli.Context) Config {
	config := Config{
		EncoderConfig: encoding.ReadCLIConfig(ctx),
		LoggerConfig:  logging.ReadCLIConfig(ctx, flags.FlagPrefix),
		ServerConfig: &encoder.ServerConfig{
			GrpcPort:              ctx.GlobalString(flags.GrpcPortFlag.Name),
			MaxConcurrentRequests: ctx.GlobalInt(flags.MaxConcurrentRequestsFlag.Name),
			RequestPoolSize:       ctx.GlobalInt(flags.RequestPoolSizeFlag.Name),
		},
		MetricsConfig: encoder.MetrisConfig{
			HTTPPort:      ctx.GlobalString(flags.MetricsHTTPPort.Name),
			EnableMetrics: ctx.GlobalBool(flags.EnableMetrics.Name),
		},
	}
	return config
}
