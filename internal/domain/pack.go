package domain

// Pack represents a single pack
type Pack struct {
	UID  string
	Size int
}

// PackCalculateRequest represents a request to calculate pack combinations
type PackCalculateRequest struct {
	Items int
}

// PackCalculateResponse represents a response to a pack calculation
type PackCalculateResponse struct {
	Calculations []*Calculation
}

// PackGetAllRequest represents a request to get all packs
type PackGetAllRequest struct {
	Page  int
	Limit int
}

// PackGetAllResponse represents a response to a pack get all request
type PackGetAllResponse struct {
	Packs []*Pack
}

// PackCreateRequest represents a request to create a new pack
type PackCreateRequest struct {
	UID  string
	Size int
}

// PackCreateResponse represents a response to a pack create request
type PackCreateResponse struct {
	Pack *Pack
}

// PackUpdateRequest represents a request to update an existing pack
type PackUpdateRequest struct {
	UID  string
	Size int
}

// PackUpdateResponse represents a response to a pack update request
type PackUpdateResponse struct {
	Pack *Pack
}

// PackDeleteRequest represents a request to delete an existing pack
type PackDeleteRequest struct {
	UID string
}

// PackDeleteResponse represents a response to a pack delete request
type PackDeleteResponse struct {
	Pack *Pack
}
