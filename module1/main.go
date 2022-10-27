package main

import (
	"github.com/ipfs/go-ipfs/plugin/plugins/git"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	_ = git.Plugins
	logger.Info("Hello world")
}
