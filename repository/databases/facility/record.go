package facility

import (
	"github.com/sewakantor/sw-be/businesses/facility"
	"gorm.io/gorm"
)

type Facility struct {
	gorm.Model
	Name       string `gorm:"type:varchar(255)"`
	Lat        float64 `gorm:"type:decimal(10,8)"`
	Lng        float64 `gorm:"type:decimal(11,8)"`
}

type Result struct {
	Distance   float64
	Latitude   float64
	Longitude  float64
	Name       string
}

func fromDomainFacility(data *facility.Domain) *Facility {
	return &Facility{
		Name:       data.Name,
		Lat:        data.Latitude,
		Lng:        data.Longitude,
	}
}

func toDomainFacility(data *Facility) *facility.Domain {
	return &facility.Domain{
		ID:         data.ID,
		Name:       data.Name,
		Latitude:   data.Lat,
		Longitude:  data.Lng,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}
}

func resultToDomain(data Result) facility.Domain {
	return facility.Domain{
		Name:       data.Name,
		Latitude:   data.Latitude,
		Longitude:  data.Longitude,
	}
}

func resultsToDomain(data []Result) []facility.Domain {
	var res []facility.Domain
	for _, s := range data {
		if s.Distance < 20 {
			res = append(res, resultToDomain(s))
		}
	}
	return res
}