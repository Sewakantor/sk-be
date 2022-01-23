package property

import (
	"github.com/sewakantor/sw-be/businesses/property"
	"gorm.io/gorm"
	"time"
)

type Complex struct {
	gorm.Model
	Name       string `gorm:"type:varchar(255)"`
	Street     string `gorm:"type:varchar(255)"`
	City       string `gorm:"type:varchar(255)"`
	State      string `gorm:"type:varchar(255)"`
	Country    string `gorm:"type:varchar(255)"`
	PostalCode int
	Lat        float64 `gorm:"type:decimal(10,8)"`
	Lng        float64 `gorm:"type:decimal(11,8)"`
}

type Building struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255)"`
	ComplexID    uint
	Complexes    Complex `gorm:"foreignKey:ComplexID"`
	Photo        string  `gorm:"type:varchar(255)"`
	Year         int
	Floor        int
	FloorSurface int
	TotalSurface int
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func fromDomainComplex(data *property.Complex) *Complex {
	return &Complex{
		Name:       data.Name,
		Street:     data.Street,
		City:       data.City,
		State:      data.State,
		Country:    data.Country,
		PostalCode: data.PostalCode,
		Lat:        data.Latitude,
		Lng:        data.Longitude,
	}
}

func toDomainComplex(data *Complex) *property.Complex {
	return &property.Complex{
		ID:         data.ID,
		Name:       data.Name,
		Street:     data.Street,
		City:       data.City,
		State:      data.State,
		Country:    data.Country,
		PostalCode: data.PostalCode,
		Latitude:   data.Lat,
		Longitude:  data.Lng,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}
}

func ToComplexesDomain(data []Complex) []property.Complex {
	var res []property.Complex
	for _, s := range data {
		res = append(res, *toDomainComplex(&s))
	}
	return res
}

func fromDomainBuilding(data *property.Building) *Building {
	return &Building{
		Name:         data.Name,
		ComplexID:    data.ComplexID,
		Photo:        data.Photo,
		Year:         data.Year,
		Floor:        data.Floor,
		FloorSurface: data.FloorSurface,
		TotalSurface: data.TotalSurface,
		Price:        data.Price,
	}
}

func toDomainBuilding(data *Building) *property.Building {
	return &property.Building{
		ID:           data.ID,
		Name:         data.Name,
		ComplexID:    data.ComplexID,
		Photo:        data.Photo,
		Year:         data.Year,
		Floor:        data.Floor,
		FloorSurface: data.FloorSurface,
		TotalSurface: data.TotalSurface,
		Price:        data.Price,
		CreatedAt:    data.CreatedAt,
		Complexes:    *toDomainComplex(&data.Complexes),
		UpdatedAt:    data.UpdatedAt,
	}
}

func ToBuildingsDomain(data []Building) []property.Building {
	var res []property.Building
	for _, s := range data {
		if s.Complexes.ID != 0 {
			res = append(res, *toDomainBuilding(&s))
		}
	}
	return res
}
