package models

type Pet struct {
	ID          int `gorm:"PrimaryKey"`
	SpeciesName string
	Species     *Species `gorm:"foreignKey:SpeciesName"`
	Happiness   int
	Energy      int
	Mood        Mood
	Exp         int
	Money       int
	HouseName   string
	House       *House `gorm:"foreignKey:HouseName"`
	// TODO: Filth int
}

type Mood int

const (
	UNDEFINED Mood = iota
	IDLE
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
