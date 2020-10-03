package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
)

var _ state = &playing{}

type playing struct{}

func (p *playing) Tick(ms int) {}

func (p *playing) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (p *playing) ReceiveMouseEvent(event interaction.MouseEvent) {}
