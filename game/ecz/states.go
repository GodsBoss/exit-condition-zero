package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

func States(spriteMap sprite.Map) map[string]game.State {
	lvls := levelsData()

	return map[string]game.State{
		"game_over":    game.ToState(newGameOver(spriteMap)),
		"level_select": game.ToState(newLevelSelect(spriteMap, lvls)),
		"playing":      game.ToState(newPlaying(spriteMap, lvls)),
		"title":        game.ToState(newTitle(spriteMap)),
	}
}
