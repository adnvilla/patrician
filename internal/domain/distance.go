package domain

// Distance represents a distance between two cities measured in days.
type Distance struct {
	Entity
	FromCity string
	ToCity   string
	Value    float32
}
