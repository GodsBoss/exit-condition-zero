package main

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/ld46/pkg/console"

	"github.com/GodsBoss/gggg/pkg/dominit"
)

func main() {
	if err := run(); err != nil {
		console.Global().LogMessage("Could not run game: %s", err.Error())
	}
}

func run() error {
	dominit.Run(game.New())
	<-make(chan struct{}, 0)
	return nil
}
