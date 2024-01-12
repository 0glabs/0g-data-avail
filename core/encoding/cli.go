package encoding

import (
	"runtime"

	"github.com/urfave/cli"
	"github.com/zero-gravity-labs/zerog-data-avail/common"
	"github.com/zero-gravity-labs/zerog-data-avail/pkg/encoding/kzgEncoder"
)

const (
	G1PathFlagName            = "kzg.g1-path"
	G2PathFlagName            = "kzg.g2-path"
	CachePathFlagName         = "kzg.cache-path"
	SRSOrderFlagName          = "kzg.srs-order"
	NumWorkerFlagName         = "kzg.num-workers"
	VerboseFlagName           = "kzg.verbose"
	PreloadEncoderFlagName    = "kzg.preload-encoder"
	CacheEncodedBlobsFlagName = "cache-encoded-blobs"
)

func CLIFlags(envPrefix string) []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:     G1PathFlagName,
			Usage:    "Path to G1 SRS",
			Required: true,
			EnvVar:   common.PrefixEnvVar(envPrefix, "G1_PATH"),
		},
		cli.StringFlag{
			Name:     G2PathFlagName,
			Usage:    "Path to G2 SRS",
			Required: true,
			EnvVar:   common.PrefixEnvVar(envPrefix, "G2_PATH"),
		},
		cli.StringFlag{
			Name:     CachePathFlagName,
			Usage:    "Path to SRS Table directory",
			Required: true,
			EnvVar:   common.PrefixEnvVar(envPrefix, "CACHE_PATH"),
		},
		cli.Uint64Flag{
			Name:     SRSOrderFlagName,
			Usage:    "Order of the SRS",
			Required: true,
			EnvVar:   common.PrefixEnvVar(envPrefix, "SRS_ORDER"),
		},
		cli.Uint64Flag{
			Name:     NumWorkerFlagName,
			Usage:    "Number of workers for multithreading",
			Required: false,
			EnvVar:   common.PrefixEnvVar(envPrefix, "NUM_WORKERS"),
			Value:    uint64(runtime.GOMAXPROCS(0)),
		},
		cli.BoolFlag{
			Name:     VerboseFlagName,
			Usage:    "Enable to see verbose output for encoding/decoding",
			Required: false,
			EnvVar:   common.PrefixEnvVar(envPrefix, "VERBOSE"),
		},
		cli.BoolFlag{
			Name:     CacheEncodedBlobsFlagName,
			Usage:    "Enable to cache encoded results",
			Required: false,
			EnvVar:   common.PrefixEnvVar(envPrefix, "CACHE_ENCODED_BLOBS"),
		},
		cli.BoolFlag{
			Name:     PreloadEncoderFlagName,
			Usage:    "Set to enable Encoder PreLoading",
			Required: false,
			EnvVar:   common.PrefixEnvVar(envPrefix, "PRELOAD_ENCODER"),
		},
	}
}

func ReadCLIConfig(ctx *cli.Context) EncoderConfig {
	cfg := kzgEncoder.KzgConfig{}
	cfg.G1Path = ctx.GlobalString(G1PathFlagName)
	cfg.G2Path = ctx.GlobalString(G2PathFlagName)
	cfg.CacheDir = ctx.GlobalString(CachePathFlagName)
	cfg.SRSOrder = ctx.GlobalUint64(SRSOrderFlagName)
	cfg.NumWorker = ctx.GlobalUint64(NumWorkerFlagName)
	cfg.Verbose = ctx.GlobalBool(VerboseFlagName)
	cfg.PreloadEncoder = ctx.GlobalBool(PreloadEncoderFlagName)
	return EncoderConfig{
		KzgConfig:         cfg,
		CacheEncodedBlobs: ctx.GlobalBoolT(CacheEncodedBlobsFlagName),
	}
}
