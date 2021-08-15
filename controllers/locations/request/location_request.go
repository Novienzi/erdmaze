package request

import (
	"erdmaze/businesses/locations"
)

type Locations struct {
	Name string `json:"name"`
}

func (req *Locations) ToDomain() *locations.Domain {
	return &locations.Domain{
		Name: req.Name,
	}
}
