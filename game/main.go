package main

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/imageload"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
	"github.com/GodsBoss/ld46/pkg/console"

	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/dominit"
)

func main() {
	if err := run(); err != nil {
		console.Global().LogMessage("Could not run game: %s", err.Error())
	}
}

func run() error {
	win, err := dom.GlobalWindow()
	if err != nil {
		return err
	}
	doc, err := win.Document()
	if err != nil {
		return err
	}
	loader := imageload.NewLoader(doc)
	img, err := loader.Load("gfx.png")
	if err != nil {
		return err
	}
	dominit.Run(game.New(sprite.NewMap(img, spriteSources)))
	<-make(chan struct{}, 0)
	return nil
}
