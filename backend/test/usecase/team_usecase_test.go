package test

import (
	"fmt"
	"project/internal/pkg/models"
	"project/internal/pkg/models/mocks"
	"project/internal/pkg/team/usecase"
	testutils2 "project/test/testutils"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type TeamTestSuite struct {
	suite.Suite
	uc        models.TeamUsecaseI
	tRepoMock *mocks.TeamRepositoryI
	tBuilder  *testutils2.TeamBuilder
}

func TestTeamTestSuite(t *testing.T) {
	suite.RunSuite(t, new(TeamTestSuite))
}

func (s *TeamTestSuite) BeforeEach(t provider.T) {
	s.tRepoMock = mocks.NewTeamRepositoryI(t)
	s.uc = usecase.NewTeamUsecase(s.tRepoMock)
	s.tBuilder = testutils2.NewTeamBuilder()
}

func (s *TeamTestSuite) TestGetAll(t provider.T) {
	ids := []int{1, 2, 3}
	names := []string{"name1", "name2", "name3"}
	countries := []string{"country1", "country2", "country3"}
	bases := []string{"base1", "base2", "base3"}

	Teams := testutils2.BuildTeams(s.tBuilder, ids, names, countries, bases)

	s.tRepoMock.On("GetAll").Return(Teams, nil)
	res, err := s.uc.GetAll()

	t.Assert().NoError(err)
	t.Assert().Equal(Teams, res)
}

func (s *TeamTestSuite) TestGetTeamById(t provider.T) {
	team := s.tBuilder.WithID(1).WithName("name1").WithCountry("country1").WithBase("base").Build()

	s.tRepoMock.On("GetTeamById", team.ID).Return(&team, nil)
	res, err := s.uc.GetTeamById(team.ID)

	t.Assert().NoError(err)
	t.Assert().Equal(team, *res)
}

func (s *TeamTestSuite) TestCreate(t provider.T) {
	team := s.tBuilder.WithID(1).WithName("name1").WithCountry("country1").WithBase("base").Build()

	s.tRepoMock.On("Create", &team).Return(1, nil)
	s.tRepoMock.On("GetTeamById", team.ID).Return(&team, nil)

	id, err := s.uc.Create(&team)
	t.Assert().NoError(err)

	res, err := s.uc.GetTeamById(team.ID)
	t.Assert().NoError(err)

	t.Assert().Equal(1, id)
	t.Assert().Equal(&team, res)
}

func (s *TeamTestSuite) TestUpdate(t provider.T) {
	team := s.tBuilder.WithID(1).WithName("name2").WithCountry("country2").WithBase("date1").Build()

	s.tRepoMock.On("Update", &team).Return(nil)
	s.tRepoMock.On("GetTeamById", team.ID).Return(&team, nil)

	err := s.uc.Update(1, &team)
	t.Assert().NoError(err)

	res, err := s.uc.GetTeamById(team.ID)
	t.Assert().NoError(err)
	t.Assert().Equal(&team, res)
}

func (s *TeamTestSuite) TestGetAllError(t provider.T) {
	s.tRepoMock.On("GetAll").Return(nil, fmt.Errorf("An error"))
	res, err := s.uc.GetAll()

	t.Assert().Error(err)
	t.Assert().Nil(res)
}
