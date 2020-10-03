package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type levelSelect struct{}

func (ls *levelSelect) Tick(ms int) {}

func (ls *levelSelect) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (ls *levelSelect) ReceiveMouseEvent(event interaction.MouseEvent) {}
