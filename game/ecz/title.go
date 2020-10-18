package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

type title struct {
	spriteMap sprite.Map
}

func newTitle(spriteMap sprite.Map) *title {
	return &title{
		spriteMap: spriteMap,
	}
}

func (t *title) ReceiveMouseEvent(event interaction.MouseEvent) *game.Transition {
	if event.Type == interaction.MouseUp && event.PrimaryButton() {
		return &game.Transition{
			NextState: "level_select",
		}
	}
	return nil
}

func (t *title) Renderables(scale int) []game.Renderable {
	return []game.Renderable{
		t.spriteMap.Produce("bg_title", 0, 0, scale, 0),
	}
}
