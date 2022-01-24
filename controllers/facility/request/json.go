package request

import (
	"github.com/sewakantor/sw-be/businesses/facility"
)

type Facility struct {
	Name string  `json:"name" validate:"required"`
	Lat  float64 `json:"lat" validate:"required"`
	Lng  float64 `json:"long" validate:"required"`
}

func (rec *Facility) FacilityToDomain() *facility.Domain {
	return &facility.Domain{
		Name:      rec.Name,
		Latitude:  rec.Lat,
		Longitude: rec.Lng,
	}
}
