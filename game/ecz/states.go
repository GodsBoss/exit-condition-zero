package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

func States(spriteMap sprite.Map) map[string]game.State {
	return map[string]game.State{
		"game_over":    NewGameOver(spriteMap),
		"level_select": NewLevelSelect(spriteMap),
		"playing":      NewPlaying(spriteMap),
		"title":        NewTitle(spriteMap),
	}
}
