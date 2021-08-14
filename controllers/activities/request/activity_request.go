package request

import (
	"erdmaze/businesses/activities"
)

type Activities struct {
	Name string `json:"name"`
}

func (req *Activities) ToDomain() *activities.Domain {
	return &activities.Domain{
		Name: req.Name,
	}
}
