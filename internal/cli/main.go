package cli

import (
	"context"
	"github.com/GeoDB-Limited/odin-evm-bridge/internal/config"
	"github.com/GeoDB-Limited/odin-evm-bridge/internal/services/deployer"
	"github.com/alecthomas/kingpin"
	"os"
)

func Run(args []string) bool {
	cfg := config.NewConfig(os.Getenv("CONFIG"))
	ctx := context.Background()
	log := cfg.Logger()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.Error("app panicked\n", rvr)
		}
	}()

	app := kingpin.New("odin-evm-bridge", "")

	runCmd := app.Command("run", "run command")
	deployService := runCmd.Command("deployer", "run a service to deploy a bridge contract")

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	switch cmd {
	case deployService.FullCommand():
		if err := deployer.New(ctx, cfg).Run(); err != nil {
			log.WithError(err).Error("failed to run deployer service")
			return false
		}
	default:
		log.WithField("command", cmd).Error("Unknown command")
		return false
	}

	return true
}
