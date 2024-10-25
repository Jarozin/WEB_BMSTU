package test

import (
	"project/internal/pkg/driver/usecase"
	"project/internal/pkg/models"
	"project/internal/pkg/models/mocks"
	testutils2 "project/test/testutils"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
)

type DriverTestSuite struct {
	suite.Suite
	uc        models.DriverUsecaseI
	dRepoMock *mocks.DriverRepositoryI
	dBuilder  *testutils2.DriverBuilder
}

func TestDriverTestSuite(t *testing.T) {
	suite.RunSuite(t, new(DriverTestSuite))
}

func (s *DriverTestSuite) BeforeEach(t provider.T) {
	s.dRepoMock = mocks.NewDriverRepositoryI(t)
	s.uc = usecase.NewDriverUsecase(s.dRepoMock)
	s.dBuilder = testutils2.NewDriverBuilder()
}

func (s *DriverTestSuite) TestGetAll(t provider.T) {
	ids := []int{1, 2, 3}
	names := []string{"name1", "name2", "name3"}
	countries := []string{"country1", "country2", "country3"}
	dates := []string{"date1", "date2", "date3"}

	drivers := testutils2.BuildDrivers(s.dBuilder, ids, names, countries, dates)

	s.dRepoMock.On("GetAll").Return(drivers, nil)
	res, err := s.uc.GetAll()

	t.Assert().NoError(err)
	t.Assert().Equal(drivers, res)
}

func (s *DriverTestSuite) TestGetDriverById(t provider.T) {
	driver := s.dBuilder.WithID(1).WithName("name1").WithCountry("country1").WithBirthDate("date1").Build()

	s.dRepoMock.On("GetDriverById", driver.ID).Return(&driver, nil)
	res, err := s.uc.GetDriverById(driver.ID)

	t.Assert().NoError(err)
	t.Assert().Equal(driver, *res)
}

func (s *DriverTestSuite) TestCreate(t provider.T) {
	driver := s.dBuilder.WithID(1).WithName("name1").WithCountry("country1").WithBirthDate("date1").Build()

	s.dRepoMock.On("Create", &driver).Return(1, nil)
	s.dRepoMock.On("GetDriverById", driver.ID).Return(&driver, nil)

	id, err := s.uc.Create(&driver)
	t.Assert().NoError(err)

	res, err := s.uc.GetDriverById(driver.ID)
	t.Assert().NoError(err)

	t.Assert().Equal(1, id)
	t.Assert().Equal(&driver, res)
}

func (s *DriverTestSuite) TestUpdate(t provider.T) {
	driverUpd := s.dBuilder.WithID(1).WithName("name2").WithCountry("country2").WithBirthDate("date1").Build()

	s.dRepoMock.On("Update", &driverUpd).Return(nil)
	s.dRepoMock.On("GetDriverById", driverUpd.ID).Return(&driverUpd, nil)

	err := s.uc.Update(1, &driverUpd)
	t.Assert().NoError(err)

	res, err := s.uc.GetDriverById(driverUpd.ID)
	t.Assert().NoError(err)
	t.Assert().Equal(&driverUpd, res)
}

func (s *DriverTestSuite) TestGetDriverByIdError(t provider.T) {
	driver := s.dBuilder.WithID(1).WithName("name1").WithCountry("country1").WithBirthDate("date1").Build()

	s.dRepoMock.On("GetDriverById", driver.ID).Return(nil, assert.AnError)
	res, err := s.uc.GetDriverById(driver.ID)

	t.Assert().Error(err)
	t.Assert().Nil(res)
}
