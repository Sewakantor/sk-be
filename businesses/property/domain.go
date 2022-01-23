package property

import (
	"github.com/sewakantor/sw-be/businesses/users"
	"time"
)

type Complex struct {
	ID         uint
	Name       string
	Street     string
	City       string
	State      string
	Country    string
	PostalCode int
	Latitude   float64
	Longitude  float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Building struct {
	ID           uint
	Name         string
	ComplexID    uint
	Complexes    Complex
	Photo        string
	Year         int
	Floor        int
	FloorSurface int
	TotalSurface int
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Review struct {
	ID         uint
	BuildingID uint
	Buildings  Building
	UserID     uint
	Users      users.Domain
	Commend    string
	Star       int
	Status     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	AddComplex(data *Complex) (*Complex, error)
	DeleteComplex(ID string) error
	GetAllComplex() ([]Complex, error)
	UpdateComplex(data *Complex, ID string) (*Complex, error)
	AddBuilding(data *Building, ID string) (*Building, error)
	GetAllBuilding(complexName string) ([]Building, error)
	GetRecommendedBuilding(limit string) ([]Building, error)
	DeleteBuilding(ID string) error
	GetSingleBuilding(ID string) (*Building, error)
	UpdateBuilding(data *Building, ID string) (*Building, error)
	AddReview(data *Review, buildingID string, usersID uint) (*Review, error)
}

type Repository interface {
	StoreComplex(data *Complex) (*Complex, error)
	DeleteComplex(ID uint64) error
	GetComplexByID(ID uint64) (*Complex, error)
	GetAllComplex() ([]Complex, error)
	GetComplexByName(name string) (*Complex, error)
	UpdateComplex(data *Complex, ID uint64) (*Complex, error)
	StoreBuilding(data *Building) (*Building, error)
	GetAllBuilding(complexName string) ([]Building, error)
	GetRecommendedBuilding(limit int) ([]Building, error)
	DeleteBuilding(ID uint) error
	GetBuildingByID(ID uint) (*Building, error)
	GetSingleBuilding(ID uint) (*Building, error)
	UpdateBuilding(data *Building, ID uint) (*Building, error)
	StoreReview(data *Review) (*Review, error)
}
