package test

import (
	"project/internal/pkg/models"
	"project/internal/pkg/models/mocks"
	"project/internal/pkg/race_result/usecase"
	testutils2 "project/test/testutils"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
)

type raceResultTestSuite struct {
	suite.Suite
	uc        models.RaceResultUsecaseI
	rRepoMock *mocks.RaceResultRepositoryI
	rBuilder  *testutils2.RaceResBuilder
}

func TestRaceResultsTestSuite(t *testing.T) {
	suite.RunSuite(t, new(raceResultTestSuite))
}

func (s *raceResultTestSuite) BeforeEach(t provider.T) {
	s.rRepoMock = mocks.NewRaceResultRepositoryI(t)
	s.uc = usecase.NewRaceResultUsecase(s.rRepoMock)
	s.rBuilder = testutils2.NewRaceResBuilder()
}

func (s *raceResultTestSuite) TestGetraceResultsOfGP(t provider.T) {
	ids := []int{1, 2, 3}
	driverPlaces := []int{1, 2, 3}
	driverIds := []int{1, 2, 3}
	driverNames := []string{"name1", "name2", "name3"}
	teamIds := []int{11, 22, 33}
	teamNames := []string{"teamname1", "teamname2", "teamname3"}
	gpIds := []int{11, 11, 11}
	gpNames := []string{"gpname1", "gpname2", "gpname3"}

	_, raceViews := testutils2.BuildRaceResults(s.rBuilder, ids, driverPlaces, driverIds, driverNames, teamIds, teamNames, gpIds, gpNames)

	s.rRepoMock.On("GetRaceResultsOfGP", 11).Return(raceViews, nil)
	res, err := s.uc.GetRaceResultsOfGP(11)

	t.Assert().NoError(err)
	t.Assert().Equal(raceViews, res)
}

func (s *raceResultTestSuite) TestGetraceResultById(t provider.T) {
	_, raceView := s.rBuilder.
		WithID(1).
		WithDriverPlace(1).
		WithDriverId(12).
		WithDriverName("name1").
		WithTeamId(32).
		WithTeamName("nameteam1").
		WithGPId(11).
		WithGPName("gpName").
		Build()

	s.rRepoMock.On("GetRaceResultById", raceView.ID).Return(&raceView, nil)
	res, err := s.uc.GetRaceResultById(raceView.ID)

	t.Assert().NoError(err)
	t.Assert().Equal(raceView, *res)
}

func (s *raceResultTestSuite) TestCreate(t provider.T) {
	race, raceView := s.rBuilder.
		WithID(1).
		WithDriverPlace(1).
		WithDriverId(12).
		WithDriverName("name1").
		WithTeamId(32).
		WithTeamName("nameteam1").
		WithGPId(11).
		WithGPName("gpName").
		Build()

	s.rRepoMock.On("Create", &race).Return(1, nil)
	s.rRepoMock.On("GetRaceResultById", raceView.ID).Return(&raceView, nil)

	id, err := s.uc.Create(&race)
	t.Assert().NoError(err)

	res, err := s.uc.GetRaceResultById(id)
	t.Assert().NoError(err)

	t.Assert().Equal(1, id)
	t.Assert().Equal(&raceView, res)
}

func (s *raceResultTestSuite) TestUpdate(t provider.T) {
	race, _ := s.rBuilder.
		WithID(1).
		WithDriverPlace(1).
		WithDriverId(12).
		WithDriverName("name1").
		WithTeamId(32).
		WithTeamName("nameteam1").
		WithGPId(11).
		WithGPName("gpName").
		Build()

	s.rRepoMock.On("Update", &race).Return(nil)

	err := s.uc.Update(1, &race)
	t.Assert().NoError(err)
}

func (s *raceResultTestSuite) TestGetraceResultByIdError(t provider.T) {
	id := 1

	s.rRepoMock.On("GetRaceResultById", id).Return(nil, assert.AnError)
	res, err := s.uc.GetRaceResultById(id)

	t.Assert().Error(err)
	t.Assert().Nil(res)
}
