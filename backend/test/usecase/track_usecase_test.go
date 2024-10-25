package test

import (
	"project/internal/pkg/models"
	"project/internal/pkg/models/mocks"
	"project/internal/pkg/track/usecase"
	testutils2 "project/test/testutils"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type TrackTestSuite struct {
	suite.Suite
	uc        models.TrackUsecaseI
	tRepoMock *mocks.TrackRepositoryI
	tBuilder  *testutils2.TrackBuilder
}

func TestTrackTestSuite(t *testing.T) {
	suite.RunSuite(t, new(TrackTestSuite))
}

func (s *TrackTestSuite) BeforeEach(t provider.T) {
	s.tRepoMock = mocks.NewTrackRepositoryI(t)
	s.uc = usecase.NewTrackUsecase(s.tRepoMock)
	s.tBuilder = testutils2.NewTrackBuilder()
}

func (s *TrackTestSuite) TestGetAll(t provider.T) {
	ids := []int{1, 2, 3}
	names := []string{"name1", "name2", "name3"}
	countries := []string{"country1", "country2", "country3"}
	towns := []string{"town1", "town2", "town3"}

	Tracks := testutils2.BuildTracks(s.tBuilder, ids, names, countries, towns)

	s.tRepoMock.On("GetAll").Return(Tracks, nil)
	res, err := s.uc.GetAll()

	t.Assert().NoError(err)
	t.Assert().Equal(Tracks, res)
}

func (s *TrackTestSuite) TestGetTrackById(t provider.T) {
	track := s.tBuilder.WithID(1).WithName("name1").WithCountry("country1").WithTown("town").Build()

	s.tRepoMock.On("GetTrackById", track.ID).Return(&track, nil)
	res, err := s.uc.GetTrackById(track.ID)

	t.Assert().NoError(err)
	t.Assert().Equal(track, *res)
}

func (s *TrackTestSuite) TestCreate(t provider.T) {
	track := s.tBuilder.WithID(1).WithName("name1").WithCountry("country1").WithTown("town").Build()

	s.tRepoMock.On("Create", &track).Return(1, nil)
	s.tRepoMock.On("GetTrackById", track.ID).Return(&track, nil)

	id, err := s.uc.Create(&track)
	t.Assert().NoError(err)

	res, err := s.uc.GetTrackById(track.ID)
	t.Assert().NoError(err)

	t.Assert().Equal(1, id)
	t.Assert().Equal(&track, res)
}

func (s *TrackTestSuite) TestUpdate(t provider.T) {
	track := s.tBuilder.WithID(1).WithName("name2").WithCountry("country2").WithTown("town").Build()

	s.tRepoMock.On("Update", &track).Return(nil)
	s.tRepoMock.On("GetTrackById", track.ID).Return(&track, nil)

	err := s.uc.Update(1, &track)
	t.Assert().NoError(err)

	res, err := s.uc.GetTrackById(track.ID)
	t.Assert().NoError(err)
	t.Assert().Equal(&track, res)
}
