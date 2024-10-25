package test

import (
	"project/internal/pkg/grand_prix/usecase"
	"project/internal/pkg/models"
	"project/internal/pkg/models/mocks"
	testutils2 "project/test/testutils"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
)

type GPTestSuite struct {
	suite.Suite
	uc         models.GrandPrixUsecaseI
	gpRepoMock *mocks.GrandPrixRepositoryI
	gpBuilder  *testutils2.GPBuilder
}

func TestGPTestSuite(t *testing.T) {
	suite.RunSuite(t, new(GPTestSuite))
}

func (s *GPTestSuite) BeforeEach(t provider.T) {
	s.gpRepoMock = mocks.NewGrandPrixRepositoryI(t)
	s.uc = usecase.NewGrandPrixUsecase(s.gpRepoMock)
	s.gpBuilder = testutils2.NewGPBuilder()
}

func (s *GPTestSuite) TestGetAll(t provider.T) {
	ids := []int{1, 2, 3}
	seasons := []int{2021, 2022, 2023}
	names := []string{"name1", "name2", "name3"}
	dateNums := []int{11, 22, 14}
	months := []string{"may", "april", "july"}
	places := []string{"place1", "place2", "place3"}
	trackIds := []int{100, 20, 14}

	gps := testutils2.BuildGrandPrixes(s.gpBuilder, ids, seasons, names, dateNums, months, places, trackIds)

	s.gpRepoMock.On("GetAll").Return(gps, nil)
	res, err := s.uc.GetAll()

	t.Assert().NoError(err)
	t.Assert().Equal(gps, res)
}

func (s *GPTestSuite) TestGetDriverById(t provider.T) {
	gp := s.gpBuilder.
		WithID(1).
		WithSeason(2021).
		WithName("name1").
		WithDateNum(12).
		WithMonth("may").
		WithPlace("place1").
		WithTrackId(11).
		Build()

	s.gpRepoMock.On("GetGPById", gp.ID).Return(&gp, nil)
	res, err := s.uc.GetGPById(gp.ID)

	t.Assert().NoError(err)
	t.Assert().Equal(gp, *res)
}

func (s *GPTestSuite) TestCreate(t provider.T) {
	gp := s.gpBuilder.
		WithID(1).
		WithSeason(2021).
		WithName("name1").
		WithDateNum(12).
		WithMonth("may").
		WithPlace("place1").
		WithTrackId(11).
		Build()

	s.gpRepoMock.On("Create", &gp).Return(1, nil)
	s.gpRepoMock.On("GetGPById", gp.ID).Return(&gp, nil)

	id, err := s.uc.Create(&gp)
	t.Assert().NoError(err)

	res, err := s.uc.GetGPById(gp.ID)
	t.Assert().NoError(err)

	t.Assert().Equal(1, id)
	t.Assert().Equal(&gp, res)
}

func (s *GPTestSuite) TestUpdate(t provider.T) {
	gp := s.gpBuilder.
		WithID(1).
		WithSeason(2021).
		WithName("name1").
		WithDateNum(12).
		WithMonth("may").
		WithPlace("place1").
		WithTrackId(11).
		Build()

	s.gpRepoMock.On("Update", &gp).Return(nil)
	s.gpRepoMock.On("GetGPById", gp.ID).Return(&gp, nil)

	err := s.uc.Update(1, &gp)
	t.Assert().NoError(err)

	res, err := s.uc.GetGPById(gp.ID)
	t.Assert().NoError(err)
	t.Assert().Equal(&gp, res)
}

func (s *GPTestSuite) TestGetAllError(t provider.T) {
	// ids := []int{1, 2, 3}
	// seasons := []int{2021, 2022, 2023}
	// names := []string{"name1", "name2", "name3"}
	// dateNums := []int{11, 22, 14}
	// months := []string{"may", "april", "july"}
	// places := []string{"place1", "place2", "place3"}
	// trackIds := []int{100, 20, 14}

	// _ := testutils2.BuildGrandPrixes(s.gpBuilder, ids, seasons, names, dateNums, months, places, trackIds)

	s.gpRepoMock.On("GetAll").Return(nil, assert.AnError)
	res, err := s.uc.GetAll()

	t.Assert().Error(err)
	t.Assert().Nil(res)
}
