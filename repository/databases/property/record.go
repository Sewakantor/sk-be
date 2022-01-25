package property

import (
	"github.com/sewakantor/sw-be/businesses/property"
	users2 "github.com/sewakantor/sw-be/repository/databases/users"
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
	Desc         string `gorm:"type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Unit struct {
	gorm.Model
	Name       string `gorm:"type:varchar(255)"`
	BuildingID uint
	Buildings  Building `gorm:"foreignKey:BuildingID"`
	Surface    int
	Capacity   int
}

type Review struct {
	gorm.Model
	BuildingID uint
	Buildings  Building `gorm:"foreignKey:BuildingID"`
	UserID     uint
	Users      users2.Users `gorm:"foreignKey:UserID"`
	Commend    string
	Star       int
	Status     int
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
		Desc:         data.Desc,
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
		Desc:         data.Desc,
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

func fromDomainReview(data *property.Review) *Review {
	return &Review{
		UserID:     data.UserID,
		BuildingID: data.BuildingID,
		Status:     0,
		Commend:    data.Commend,
		Star:       data.Star,
	}
}

func toDomainReview(data *Review) *property.Review {
	return &property.Review{
		ID:         data.ID,
		UserID:     data.UserID,
		BuildingID: data.BuildingID,
		Status:     data.Status,
		Commend:    data.Commend,
		Star:       data.Star,
		Buildings:  *toDomainBuilding(&data.Buildings),
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		Users:      *users2.ToDomainUser(&data.Users),
	}
}

func ToReviewsDomain(data []Review) []property.Review {
	var res []property.Review
	for _, s := range data {
		res = append(res, *toDomainReview(&s))
	}
	return res
}

func fromDomainUnit(data *property.Unit) *Unit {
	return &Unit{
		Name:       data.Name,
		Capacity:   data.Capacity,
		Surface:    data.Surface,
		BuildingID: data.BuildingID,
	}
}

func ToDomainUnit(data *Unit) *property.Unit {

	return &property.Unit{
		ID:         data.ID,
		Name:       data.Name,
		BuildingID: data.BuildingID,
		Surface:    data.Surface,
		Capacity:   data.Capacity,
		Buildings:  *toDomainBuilding(&data.Buildings),
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}
}

func ToUnitsDomain(data []Unit) []property.Unit {
	var res []property.Unit
	for _, s := range data {
		res = append(res, *ToDomainUnit(&s))
	}
	return res
}
