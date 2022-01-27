package property_test

import (
	"errors"
	"github.com/sewakantor/sw-be/businesses/property"
	_propertyMock "github.com/sewakantor/sw-be/businesses/property/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	mockPropertyRepository _propertyMock.Repository
	propertyService        property.Service

	domainTest property.Complex
)

func TestMain(m *testing.M) {
	propertyService = property.NewPropertyService(&mockPropertyRepository)
	domainTest = property.Complex{
		Name: "Area Mandiri",
	}
	m.Run()
}

func TestPropertyService_AddComplex(t *testing.T) {
	t.Run("exist", func(t *testing.T) {
		mockPropertyRepository.On("GetComplexByName", mock.Anything, mock.Anything).Return(&domainTest, nil).Once()
		domainTest = property.Complex{
			Name: "Area Mandiri",
		}

		res, err := propertyService.AddComplex(&domainTest)
		assert.NotNil(t, err)
		assert.Nil(t, res)
		//assert.Equal(t, res.Email, domainTest.Email)
	})

	t.Run("Success", func(t *testing.T) {
		mockPropertyRepository.On("GetComplexByName", mock.Anything, mock.Anything).Return(nil, errors.New("not found ...")).Once()
		mockPropertyRepository.On("StoreComplex", mock.Anything, mock.Anything).Return(&domainTest, nil).Once()
		domainTest = property.Complex{
			Name: "Area Mandiri",
		}

		res, err := propertyService.AddComplex(&domainTest)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		//assert.Equal(t, res.Email, domainTest.Email)
	})
}

func TestPropertyService_DeleteComplex(t *testing.T) {
	t.Run("error parser", func(t *testing.T) {
		err := propertyService.DeleteComplex("wadaw")
		assert.NotNil(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		mockPropertyRepository.On("GetComplexByID", mock.Anything, mock.Anything).Return(nil, errors.New("not found ...")).Once()
		domainTest = property.Complex{
			Name: "Area Mandiri",
		}

		err := propertyService.DeleteComplex("1")
		assert.NotNil(t, err)
		//assert.Equal(t, res.Email, domainTest.Email)
	})

	t.Run("error", func(t *testing.T) {
		mockPropertyRepository.On("GetComplexByID", mock.Anything, mock.Anything).Return(nil, nil).Once()
		mockPropertyRepository.On("DeleteComplex", mock.Anything, mock.Anything).Return(nil, errors.New("dwad")).Once()
		domainTest = property.Complex{
			Name: "Area Mandiri",
		}

		err := propertyService.DeleteComplex("1")
		assert.Nil(t, err)
		//assert.Equal(t, res.Email, domainTest.Email)
	})
}

func TestPropertyService_GetAllComplex(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		mockPropertyRepository.On("GetAllComplex", mock.Anything, mock.Anything).Return(nil, errors.New("not found ...")).Once()
		domainTest = property.Complex{
			Name: "Area Mandiri",
		}

		_, err := propertyService.GetAllComplex("adw")
		assert.NotNil(t, err)
		//assert.Equal(t, res.Email, domainTest.Email)
	})

	t.Run("success", func(t *testing.T) {
		mockPropertyRepository.On("GetAllComplex", mock.Anything, mock.Anything).Return(nil, nil).Once()
		domainTest = property.Complex{
			Name: "Area Mandiri",
		}

		_, err := propertyService.GetAllComplex("adw")
		assert.Nil(t, err)
		//assert.Equal(t, res.Email, domainTest.Email)
	})
}

func TestPropertyService_UpdateComplex(t *testing.T) {
	t.Run("error parser", func(t *testing.T) {
		_, err := propertyService.UpdateComplex(&property.Complex{}, "wadaw")
		assert.NotNil(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		mockPropertyRepository.On("UpdateComplex", mock.Anything, mock.Anything).Return(nil, errors.New("not found ...")).Once()
		domainTest = property.Complex{
			Name: "Area Mandiri",
		}

		_, err := propertyService.UpdateComplex(&domainTest, "121")
		assert.NotNil(t, err)
		//assert.Equal(t, res.Email, domainTest.Email)
	})

	t.Run("err", func(t *testing.T) {
		mockPropertyRepository.On("UpdateComplex", mock.Anything, mock.Anything).Return(nil, errors.New("dwaund ...")).Once()
		domainTest = property.Complex{
			Name: "Area Mandiri",
		}

		_, err := propertyService.UpdateComplex(&domainTest, "121")
		assert.NotNil(t, err)
		//assert.Equal(t, res.Email, domainTest.Email)
	})

	t.Run("success", func(t *testing.T) {
		mockPropertyRepository.On("UpdateComplex", mock.Anything, mock.Anything).Return(nil, nil).Once()
		domainTest = property.Complex{
			Name: "Area Mandiri",
		}

		_, err := propertyService.UpdateComplex(&domainTest, "121")
		assert.Nil(t, err)
		//assert.Equal(t, res.Email, domainTest.Email)
	})
}

func TestPropertyService_AddBuilding(t *testing.T) {
	mockPropertyRepository.On("StoreBuilding", mock.Anything, mock.Anything).Return(nil, errors.New("sadw")).Once()
	t.Run("error", func(t *testing.T) {
		_, err := propertyService.AddBuilding(&property.Building{}, "qs")
		assert.NotNil(t, err)
	})

	mockPropertyRepository.On("StoreBuilding", mock.Anything, mock.Anything).Return(nil, errors.New("violates foreign")).Once()
	t.Run("error not found", func(t *testing.T) {
		_, err := propertyService.AddBuilding(&property.Building{}, "qs")
		assert.NotNil(t, err)
	})

	mockPropertyRepository.On("StoreBuilding", mock.Anything, mock.Anything).Return(nil, nil).Once()
	t.Run("success", func(t *testing.T) {
		_, err := propertyService.AddBuilding(&property.Building{}, "qs")
		assert.Nil(t, err)
	})
}

func TestPropertyService_GetAllBuilding(t *testing.T) {
	mockPropertyRepository.On("GetAllBuilding", mock.Anything, mock.Anything).Return(nil, errors.New("sadw")).Once()
	t.Run("error", func(t *testing.T) {
		_, err := propertyService.GetAllBuilding("qs")
		assert.NotNil(t, err)
	})

	mockPropertyRepository.On("GetAllBuilding", mock.Anything, mock.Anything).Return(nil, nil).Once()
	t.Run("error", func(t *testing.T) {
		_, err := propertyService.GetAllBuilding("qs")
		assert.NotNil(t, err)
	})

	mockPropertyRepository.On("GetAllBuilding", mock.Anything, mock.Anything).Return([]property.Building{}, nil).Once()
	t.Run("error", func(t *testing.T) {
		_, err := propertyService.GetAllBuilding("123")
		assert.Nil(t, err)
	})
}

func TestPropertyService_GetRecommendedBuilding(t *testing.T) {
	mockPropertyRepository.On("GetRecommendedBuilding", mock.Anything, mock.Anything).Return(nil, errors.New("sadw")).Once()
	t.Run("error", func(t *testing.T) {
		_, err := propertyService.GetRecommendedBuilding("2")
		assert.NotNil(t, err)
	})

	mockPropertyRepository.On("GetRecommendedBuilding", mock.Anything, mock.Anything).Return([]property.Building{}, nil).Once()
	t.Run("error", func(t *testing.T) {
		_, err := propertyService.GetRecommendedBuilding("2")
		assert.Nil(t, err)
	})
}

func TestPropertyService_DeleteBuilding(t *testing.T) {
	t.Run("error parser", func(t *testing.T) {
		err := propertyService.DeleteBuilding("wadaw")
		assert.NotNil(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		mockPropertyRepository.On("GetBuildingByID", mock.Anything, mock.Anything).Return(nil, errors.New("not found")).Once()
		err := propertyService.DeleteBuilding("2")
		assert.NotNil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockPropertyRepository.On("GetBuildingByID", mock.Anything, mock.Anything).Return(nil, nil).Once()
		mockPropertyRepository.On("DeleteBuilding", mock.Anything, mock.Anything).Return(errors.New("d")).Once()
		err := propertyService.DeleteBuilding("2")
		assert.NotNil(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mockPropertyRepository.On("GetBuildingByID", mock.Anything, mock.Anything).Return(nil, nil).Once()
		mockPropertyRepository.On("DeleteBuilding", mock.Anything, mock.Anything).Return(nil, nil).Once()
		err := propertyService.DeleteBuilding("2")
		assert.Nil(t, err)
	})
}

func TestPropertyService_GetAllUnit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockPropertyRepository.On("GetAllUnit", mock.Anything, mock.Anything).Return(nil, errors.New("dwada")).Once()
		_, err := propertyService.GetAllUnit()
		assert.NotNil(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mockPropertyRepository.On("GetAllUnit", mock.Anything, mock.Anything).Return([]property.Unit{}, nil).Once()
		_, err := propertyService.GetAllUnit()
		assert.Nil(t, err)
	})
}

func TestPropertyService_DeleteUnit(t *testing.T) {
	t.Run("error parser", func(t *testing.T) {
		err := propertyService.DeleteUnit("wadaw")
		assert.NotNil(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mockPropertyRepository.On("GetUnitByID", mock.Anything, mock.Anything).Return(nil, errors.New("not found")).Once()
		err := propertyService.DeleteUnit("21")
		assert.NotNil(t, err)
	})

	t.Run("err", func(t *testing.T) {
		mockPropertyRepository.On("GetUnitByID", mock.Anything, mock.Anything).Return(nil, nil).Once()
		mockPropertyRepository.On("DeleteUnit", mock.Anything, mock.Anything).Return(errors.New("not dwa")).Once()
		err := propertyService.DeleteUnit("21")
		assert.NotNil(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mockPropertyRepository.On("GetUnitByID", mock.Anything, mock.Anything).Return(nil, nil).Once()
		mockPropertyRepository.On("DeleteUnit", mock.Anything, mock.Anything).Return(nil).Once()
		err := propertyService.DeleteUnit("21")
		assert.Nil(t, err)
	})
}

func TestPropertyService_AddUnit(t *testing.T) {
	t.Run("error parser", func(t *testing.T) {
		_, err := propertyService.AddUnit(&property.Unit{},"wadaw")
		assert.NotNil(t, err)
	})

	t.Run("error not found", func(t *testing.T) {
		mockPropertyRepository.On("StoreUnit", mock.Anything, mock.Anything).Return(nil, errors.New("violates foreign")).Once()
		_, err := propertyService.AddUnit(&property.Unit{},"12")
		assert.NotNil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockPropertyRepository.On("StoreUnit", mock.Anything, mock.Anything).Return(nil, errors.New("sad")).Once()
		_, err := propertyService.AddUnit(&property.Unit{},"12")
		assert.NotNil(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mockPropertyRepository.On("StoreUnit", mock.Anything, mock.Anything).Return(nil, nil).Once()
		_, err := propertyService.AddUnit(&property.Unit{},"12")
		assert.Nil(t, err)
	})
}

func TestPropertyService_GetSingleBuilding(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		mockPropertyRepository.On("GetSingleBuilding", mock.Anything, mock.Anything).Return(nil, errors.New("dwa")).Once()
		_, err := propertyService.GetSingleBuilding("12")
		assert.NotNil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockPropertyRepository.On("GetSingleBuilding", mock.Anything, mock.Anything).Return(&property.Building{ID:0}, nil).Once()
		_, err := propertyService.GetSingleBuilding("12")
		assert.NotNil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockPropertyRepository.On("GetSingleBuilding", mock.Anything, mock.Anything).Return(&property.Building{ID:1}, nil).Once()
		_, err := propertyService.GetSingleBuilding("12")
		assert.Nil(t, err)
	})
}

func TestPropertyService_UpdateBuilding(t *testing.T) {
	t.Run("error parser", func(t *testing.T) {
		_, err := propertyService.UpdateBuilding(&property.Building{},"wadaw")
		assert.NotNil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockPropertyRepository.On("UpdateBuilding", mock.Anything, mock.Anything).Return(nil, errors.New("not found")).Once()
		_, err := propertyService.UpdateBuilding(&property.Building{}, "12")
		assert.NotNil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockPropertyRepository.On("UpdateBuilding", mock.Anything, mock.Anything).Return(nil, errors.New("w found")).Once()
		_, err := propertyService.UpdateBuilding(&property.Building{}, "12")
		assert.NotNil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockPropertyRepository.On("UpdateBuilding", mock.Anything, mock.Anything).Return(nil, nil).Once()
		_, err := propertyService.UpdateBuilding(&property.Building{}, "12")
		assert.Nil(t, err)
	})
}

func TestPropertyService_AddReview(t *testing.T) {
	t.Run("error parser", func(t *testing.T) {
		_, err := propertyService.AddReview(&property.Review{},"wadaw", 2)
		assert.NotNil(t, err)
	})

	t.Run("error not found", func(t *testing.T) {
		mockPropertyRepository.On("StoreReview", mock.Anything, mock.Anything).Return(nil, errors.New("violates foreign")).Once()
		_, err := propertyService.AddReview(&property.Review{},"2", 2)
		assert.NotNil(t, err)
	})

	t.Run("error ", func(t *testing.T) {
		mockPropertyRepository.On("StoreReview", mock.Anything, mock.Anything).Return(nil, errors.New("violates forseign")).Once()
		_, err := propertyService.AddReview(&property.Review{},"2", 2)
		assert.NotNil(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mockPropertyRepository.On("StoreReview", mock.Anything, mock.Anything).Return(nil, nil).Once()
		_, err := propertyService.AddReview(&property.Review{},"2", 2)
		assert.Nil(t, err)
	})
}

func TestPropertyService_ApproveReview(t *testing.T) {
	t.Run("error parser", func(t *testing.T) {
		_, err := propertyService.ApproveReview("w")
		assert.NotNil(t, err)
	})

	t.Run("err", func(t *testing.T) {
		mockPropertyRepository.On("ApproveReview", mock.Anything, mock.Anything).Return(nil, errors.New("not found")).Once()
		_, err := propertyService.ApproveReview("2")
		assert.NotNil(t, err)
	})

	t.Run("err", func(t *testing.T) {
		mockPropertyRepository.On("ApproveReview", mock.Anything, mock.Anything).Return(nil, errors.New("snost found")).Once()
		_, err := propertyService.ApproveReview("2")
		assert.NotNil(t, err)
	})

	t.Run("err", func(t *testing.T) {
		mockPropertyRepository.On("ApproveReview", mock.Anything, mock.Anything).Return(nil, nil).Once()
		_, err := propertyService.ApproveReview("2")
		assert.Nil(t, err)
	})
}

func TestPropertyService_GetAllReview(t *testing.T) {
	t.Run("error parser", func(t *testing.T) {
		_, err := propertyService.GetAllReview("w","w", "true")
		assert.NotNil(t, err)
	})

	t.Run("error parser", func(t *testing.T) {
		_, err := propertyService.GetAllReview("1","w", "true")
		assert.NotNil(t, err)
	})

	t.Run("error parser", func(t *testing.T) {
		_, err := propertyService.GetAllReview("1","2", "2s")
		assert.NotNil(t, err)
	})

	t.Run("error parser", func(t *testing.T) {
		_, err := propertyService.GetAllReview("1","s", "")
		assert.NotNil(t, err)
	})

	t.Run("error parser", func(t *testing.T) {
		_, err := propertyService.GetAllReview("1","", "s")
		assert.NotNil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockPropertyRepository.On("GetAllReview", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("snost found")).Once()
		_, err := propertyService.GetAllReview("1","2", "true")
		assert.NotNil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockPropertyRepository.On("GetAllReview", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()
		_, err := propertyService.GetAllReview("1","2", "true")
		assert.Nil(t, err)
	})
}