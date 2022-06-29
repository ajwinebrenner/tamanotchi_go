package mood

type Mood int

const (
	IDLE Mood = iota
	HAPPY
	TIRED
	SICK
	DEAD
)

func (m Mood) String() string {
	switch m {
	case IDLE:
		return "Idle"
	case HAPPY:
		return "Happy"
	case TIRED:
		return "Tired"
	case SICK:
		return "Sick"
	case DEAD:
		return "Dead"
	}
	return "Unknown"
}
