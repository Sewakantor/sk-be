package complex

import (
	"github.com/sewakantor/sw-be/businesses/complex"
	"gorm.io/gorm"
)

type Complex struct {
	gorm.Model
	Name       string `gorm:"type:varchar(255)"`
	Street     string `gorm:"type:varchar(255)"`
	City       string `gorm:"type:varchar(255)"`
	State      string `gorm:"type:varchar(255)"`
	Country    string `gorm:"type:varchar(255)"`
	PostalCode int
	latitude   float64
	longitude  float64
}

func fromDomain(data *complex.Domain) *Complex {
	return &Complex{
		Name: data.Name,
		Street: data.Street,
		City: data.City,
		State: data.State,
		Country: data.Country,
		PostalCode: data.PostalCode,
		latitude: data.Latitude,
		longitude: data.Longitude,
	}
}

func toDomain(data *Complex) *complex.Domain {
	return &complex.Domain{
		ID: data.ID,
		Name: data.Name,
		Street: data.Street,
		City: data.City,
		State: data.State,
		Country: data.Country,
		PostalCode: data.PostalCode,
		Latitude: data.latitude,
		Longitude: data.longitude,
		CreatedAt: data.CreatedAt,
	}
}
