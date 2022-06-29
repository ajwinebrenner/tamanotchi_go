package variety

type Variety int

const (
	UNHEALTHY Variety = iota
	HEALTHY
	HEALING
)

func (v Variety) String() string {
	switch v {
	case UNHEALTHY:
		return "Unhealthy"
	case HEALTHY:
		return "Healthy"
	case HEALING:
		return "Healing"
	}
	return "Unknown"
}
