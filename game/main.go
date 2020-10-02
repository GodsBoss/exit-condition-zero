package main

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"

	"github.com/GodsBoss/gggg/pkg/dominit"
)

func main() {
	dominit.Run(game.New())
	<-make(chan struct{}, 0)
}
