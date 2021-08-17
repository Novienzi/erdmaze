package request

import tourismpackages "erdmaze/businesses/tourism_packages"

type TourismPackages struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	TotalPrice  string `json:"total_price"`
	TotalTime   string `json:"total_time"`
	LocationID  int    `json:"location_id"`
	ActivityID  int    `json:"activity_id"`
	Address     string `json:"address"`
	AddressUrl  string `json:"address_url"`
}

func (req *TourismPackages) ToDomain() *tourismpackages.Domain {
	return &tourismpackages.Domain{
		Name:        req.Name,
		Description: req.Description,
		TotalPrice:  req.TotalPrice,
		TotalTime:   req.TotalTime,
		LocationID:  req.LocationID,
		ActivityID:  req.ActivityID,
		Address:     req.Address,
		AddressUrl:  req.AddressUrl,
	}
}
