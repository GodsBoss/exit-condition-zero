// build js,wasm

package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type Game struct {
	output *dom.Context2D
}

func New() *Game {
	return &Game{}
}

func (g *Game) TicksPerSecond() int {
	return 40
}

func (g *Game) SetOutput(ctx2d *dom.Context2D) {
	ctx2d.DisableImageSmoothing()
	g.output = ctx2d
}

func (g *Game) Render() {}

func (g *Game) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, scaleX, scaleY float64) {
	sx := (availableWidth - 20) / uiWidth
	sy := (availableHeight - 20) / uiHeight
	if sx < 1 {
		sx = 1
	}
	if sy < 1 {
		sy = 1
	}
	s := sx
	if sy < s {
		s = sy
	}
	return uiWidth * s, uiHeight * s, float64(s), float64(s)
}

func (g *Game) Tick(ms int) {}

func (g *Game) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (g *Game) ReceiveMouseEvent(event interaction.MouseEvent) {}

const (
	uiWidth  = 320
	uiHeight = 240
)
