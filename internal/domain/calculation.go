package domain

// Calculation represents a single pack calculation
type Calculation struct {
	PackUID  string
	PackSize int
	Quantity int
	Items    int
}
