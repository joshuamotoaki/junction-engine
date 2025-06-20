package main

import (
	"github.com/tigerappsorg/junction-engine/application"
	"github.com/tigerappsorg/junction-engine/config"
)

func main() {
	cfg := config.Load()
	application.Run(cfg)
}
