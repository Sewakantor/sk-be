package facility

import (
	"time"
)

type Domain struct {
	ID        uint
	Name      string
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	AddFacility(data *Domain) (*Domain, error)
	UnitsByGeo(long, lat string) ([]Domain, error)
}

type Repository interface {
	Store(data *Domain) (*Domain, error)
	GetFacilityByGeo(lat, long float64) ([]Domain, error)
}
