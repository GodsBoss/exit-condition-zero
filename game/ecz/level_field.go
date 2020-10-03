package ecz

// field is a level's field when playing.
type field interface {
	// Resets this field to its initial state as found in the level's data (for
	// pre-existing fields) or configured by the player (for fields set by the player).
	Reset()

	// ExtractOutputPulses takes all the output pulses this field has to offer.
	// This may change the field, e.g. "exhaust" pulses.
	ExtractOutputPulses() []direction

	// IsHit takes a beam from a direction. Returned directions are directly
	// converted to new pulses. The field may also change state.
	// If the first return value is false, the field is not "hit" in a meaningful
	// sense. This is logically different from being hit and sending a pulse into
	// the same direction.
	IsHit(direction) (bool, []direction)

	// IsDeletable determines wether the player can delete this field.
	IsDeletable() bool

	// IsMovable determines wether the player can move this field.
	IsMovable() bool
}

type direction string

const (
	dirUp    direction = "up"
	dirRight direction = "right"
	dirDown  direction = "down"
	dirLeft  direction = "left"
)

// fieldWithVictoryCondition is an optional interface fields can implement.
// Every field will be questioned for victory when running. If any field
// disallows a victory, the level is not over.
type fieldWithVictoryCondition interface {
	// AllowsVictory returns true if this field allows for the player to be
	// victorious in this level.
	AllowsVictory() bool
}
