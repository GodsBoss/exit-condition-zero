package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

func States(spriteMap sprite.Map) map[string]game.State {
	lvls := levelsData()

	return map[string]game.State{
		"game_over":    newGameOver(spriteMap),
		"level_select": newLevelSelect(spriteMap, lvls),
		"playing":      newPlaying(spriteMap, lvls),
		"title":        newTitle(spriteMap),
	}
}
