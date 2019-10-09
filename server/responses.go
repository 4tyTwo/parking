package server

// ErrorResponse
type ErrorResponse struct {
	Description string
}

// PlaceCarResponse
type PlaceCarResponse struct {
	PlaceCode string
}

// FreePlacesResponse
type FreePlacesResponse struct {
	FreePlaces int
}
