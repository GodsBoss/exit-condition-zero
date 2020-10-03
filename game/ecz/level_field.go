package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
	"github.com/GodsBoss/exit-condition-zero/pkg/vector/int2d"
)

// field is a level's field when playing.
type field interface {
	// Resets this field to its initial state as found in the level's data (for
	// pre-existing fields) or configured by the player (for fields set by the player).
	// This is called when the player stops the run, not if the level is reset.
	Reset()

	// ExtractOutputPulses takes all the output pulses this field has to offer.
	// This may change the field, e.g. "exhaust" pulses.
	ExtractOutputPulses() []direction

	// ImmediateHit takes a beam from a direction. Returned directions are directly
	// converted to new pulses. The field should not change state.
	// The first return parameter determines wether this field was "hit" in a
	// meaningful sense. The returned directions are handled wether the field was
	// hit or not, so if a pulse is to be continued, it should be returned.
	ImmediateHit(direction) (bool, []direction)

	// Receive is called after all pulses have ended somewhere and a pulse hit
	// this field. ImmediateHit() must have returned true for Receive() to be called.
	Receive([]direction)

	// IsDeletable determines wether the player can delete this field.
	IsDeletable() bool

	// IsMovable determines wether the player can move this field.
	IsMovable() bool

	// IsConfigurable determines wether the player can configure this field.
	IsConfigurable() bool

	// Configure changes the state of the field, caused by the initiative of the
	// player.
	Configure()

	// Renderable returns the graphical representation of a field. x and y are
	// in-game unscaled pixel coordinates.
	Renderable(x, y int, scale int) game.Renderable
}

type direction string

const (
	dirUp    direction = "up"
	dirRight direction = "right"
	dirDown  direction = "down"
	dirLeft  direction = "left"
)

func turnDirectionClockwise(dir direction) direction {
	return (map[direction]direction{
		dirUp:    dirRight,
		dirRight: dirDown,
		dirDown:  dirLeft,
		dirLeft:  dirUp,
	})[dir]
}

func turnDirectionCounterClockwise(dir direction) direction {
	return (map[direction]direction{
		dirUp:    dirLeft,
		dirRight: dirUp,
		dirDown:  dirRight,
		dirLeft:  dirDown,
	})[dir]
}

func oppositeDirection(dir direction) direction {
	return (map[direction]direction{
		dirUp:    dirDown,
		dirRight: dirLeft,
		dirDown:  dirUp,
		dirLeft:  dirRight,
	})[dir]
}

var directionVectors = map[direction]int2d.Vector{
	dirUp:    int2d.Up(),
	dirRight: int2d.Right(),
	dirDown:  int2d.Down(),
	dirLeft:  int2d.Left(),
}

var directionSpriteIDs = map[direction]string{
	dirUp:    "output_up",
	dirRight: "output_right",
	dirDown:  "output_down",
	dirLeft:  "output_left",
}

func createRenderableForDirections(spriteMap sprite.Map, directions []direction, x, y int, scale int) game.Renderable {
	r := make(rendering.Renderables, 0)
	for i := range directions {
		if id, ok := directionSpriteIDs[directions[i]]; ok {
			r = append(r, spriteMap.Produce(id, x, y, scale, 0))
		}
	}
	return r
}

// fieldWithVictoryCondition is an optional interface fields can implement.
// Every field will be questioned for victory when running. If any field
// disallows a victory, the level is not over.
type fieldWithVictoryCondition interface {
	// AllowsVictory returns true if this field allows for the player to be
	// victorious in this level.
	AllowsVictory() bool
}

// fieldFree is implemented only be empty fields.
type fieldFree interface {
	// IsFree returns wether other fields may be placed here despite the field
	// neither being deletable nor movable.
	IsFree() bool
}

func isFieldFree(f field) bool {
	if ff, ok := f.(fieldFree); ok {
		return ff.IsFree()
	}
	return false
}
