package sample_test

import (
	"errors"
	"simple-golang-api/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BrandMock struct {
	Mock mock.Mock
}

func (bm *BrandMock) GetBrandById(id int) *models.JSONBrand {
	args := bm.Mock.Called(id)

	if args.Get(0) == nil {
		return nil
	} else {
		brand := args.Get(0).(models.JSONBrand)

		return &brand
	}
}

func (bm *BrandMock) CreateBrand(name string) *models.JSONBrand {
	args := bm.Mock.Called(name)

	if args.Get(0) == nil {
		return nil
	} else {
		brand := args.Get(0).(models.JSONBrand)

		return &brand
	}
}

var bm = &BrandMock{Mock: mock.Mock{}}

func Get(id int) (*models.JSONBrand, error) {
	brand := bm.GetBrandById(id)
	if brand == nil {
		return nil, errors.New("brand Not Found")
	} else {
		return brand, nil
	}
}

func Create(name string, id int) (*models.JSONBrand, error) {
	if len(name) == 0 {
		return nil, errors.New("name cannot be empty")
	} else if len(name) < 2 {
		return nil, errors.New("name minimal 3 characters")
	}
	brand := bm.CreateBrand(name)

	if brand == nil {
		return nil, errors.New("error create new brand")
	} else {
		return brand, nil
	}
}

func TestGetBrandById_Success(t *testing.T) {
	brand := models.JSONBrand{
		Id:   1,
		Name: "Dell",
	}
	// expectation
	bm.Mock.On("GetBrandById", 1).Return(brand)

	// check expectation
	res, err := Get(1)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, brand.Id, res.Id)
	assert.Equal(t, brand.Name, res.Name)
}

func TestCreateBrandWithEmptyName_Fail(t *testing.T) {
	brand := models.JSONBrand{
		Id:   1,
		Name: "",
	}
	bm.Mock.On("CreateBrand").Return(errors.New("name cannot be empty"))

	_, err := Create(brand.Name, brand.Id)
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("name cannot be empty"), err)
}
