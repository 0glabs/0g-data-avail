package flags

import (
	"github.com/0glabs/0g-data-avail/common"
	"github.com/0glabs/0g-data-avail/common/logging"
	"github.com/0glabs/0g-data-avail/core/encoding"
	"github.com/urfave/cli"
)

const (
	FlagPrefix   = "disperser-encoder"
	envVarPrefix = "DISPERSER_ENCODER"
)

var (
	/* Required Flags */
	GrpcPortFlag = cli.StringFlag{
		Name:     common.PrefixFlag(FlagPrefix, "grpc-port"),
		Usage:    "Port at which encoder listens for grpc calls",
		Required: true,
		EnvVar:   common.PrefixEnvVar(envVarPrefix, "GRPC_PORT"),
	}
	/* Optional Flags*/
	MetricsHTTPPort = cli.StringFlag{
		Name:     common.PrefixFlag(FlagPrefix, "metrics-http-port"),
		Usage:    "the http port which the metrics prometheus server is listening",
		Required: false,
		Value:    "9100",
		EnvVar:   common.PrefixEnvVar(envVarPrefix, "METRICS_HTTP_PORT"),
	}
	EnableMetrics = cli.BoolFlag{
		Name:     common.PrefixFlag(FlagPrefix, "enable-metrics"),
		Usage:    "start metrics server",
		Required: false,
		EnvVar:   common.PrefixEnvVar(envVarPrefix, "ENABLE_METRICS"),
	}
	MaxConcurrentRequestsFlag = cli.IntFlag{
		Name:     common.PrefixFlag(FlagPrefix, "max-concurrent-requests"),
		Usage:    "maximum number of concurrent requests",
		Required: false,
		Value:    16,
		EnvVar:   common.PrefixEnvVar(envVarPrefix, "MAX_CONCURRENT_REQUESTS"),
	}
	RequestPoolSizeFlag = cli.IntFlag{
		Name:     common.PrefixFlag(FlagPrefix, "request-pool-size"),
		Usage:    "maximum number of requests in the request pool",
		Required: false,
		Value:    32,
		EnvVar:   common.PrefixEnvVar(envVarPrefix, "REQUEST_POOL_SIZE"),
	}
)

var requiredFlags = []cli.Flag{
	GrpcPortFlag,
}

var optionalFlags = []cli.Flag{
	MetricsHTTPPort,
	EnableMetrics,
	MaxConcurrentRequestsFlag,
	RequestPoolSizeFlag,
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag

func init() {
	Flags = append(requiredFlags, optionalFlags...)
	Flags = append(Flags, encoding.CLIFlags(envVarPrefix)...)
	Flags = append(Flags, logging.CLIFlags(envVarPrefix, FlagPrefix)...)
}
