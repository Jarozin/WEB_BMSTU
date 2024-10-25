package test

import (
	"fmt"
	"project/internal/pkg/models"
	"project/internal/pkg/models/mocks"
	"project/internal/pkg/qual_result/usecase"
	testutils2 "project/test/testutils"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type QualResultTestSuite struct {
	suite.Suite
	uc        models.QualResultUsecaseI
	qRepoMock *mocks.QualResultRepositoryI
	qBuilder  *testutils2.QualResBuilder
}

func TestQualResultsTestSuite(t *testing.T) {
	suite.RunSuite(t, new(QualResultTestSuite))
}

func (s *QualResultTestSuite) BeforeEach(t provider.T) {
	s.qRepoMock = mocks.NewQualResultRepositoryI(t)
	s.uc = usecase.NewQualResultUsecase(s.qRepoMock)
	s.qBuilder = testutils2.NewQualResBuilder()
}

func (s *QualResultTestSuite) TestGetQualResultsOfGP(t provider.T) {
	ids := []int{1, 2, 3}
	driverPlaces := []int{1, 2, 3}
	driverIds := []int{1, 2, 3}
	driverNames := []string{"name1", "name2", "name3"}
	teamIds := []int{11, 22, 33}
	teamNames := []string{"teamname1", "teamname2", "teamname3"}
	q1Time := []time.Time{
		time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local),
		time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local),
		time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local)}
	q2Time := []time.Time{
		time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local),
		time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local),
		time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local)}
	q3Time := []time.Time{
		time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local),
		time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local),
		time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local)}
	gpIds := []int{11, 11, 11}
	gpNames := []string{"gpname1", "gpname2", "gpname3"}

	_, qualViews := testutils2.BuildQualResults(s.qBuilder, ids, driverPlaces, driverIds, driverNames, teamIds, teamNames, q1Time, q2Time, q3Time, gpIds, gpNames)

	s.qRepoMock.On("GetQualResultsOfGP", 11).Return(qualViews, nil)
	res, err := s.uc.GetQualResultsOfGP(11)

	t.Assert().NoError(err)
	t.Assert().Equal(qualViews, res)
}

func (s *QualResultTestSuite) TestGetQualResultById(t provider.T) {
	_, qualView := s.qBuilder.
		WithID(1).
		WithDriverPlace(1).
		WithDriverId(12).
		WithDriverName("name1").
		WithTeamId(32).
		WithTeamName("nameteam1").
		WithQ1time(time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local)).
		WithQ2time(time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local)).
		WithQ3time(time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local)).
		WithGPId(11).
		WithGPName("gpName").
		Build()

	s.qRepoMock.On("GetQualResultById", qualView.ID).Return(&qualView, nil)
	res, err := s.uc.GetQualResultById(qualView.ID)

	t.Assert().NoError(err)
	t.Assert().Equal(qualView, *res)
}

func (s *QualResultTestSuite) TestCreate(t provider.T) {
	qual, qualView := s.qBuilder.
		WithID(1).
		WithDriverPlace(1).
		WithDriverId(12).
		WithDriverName("name1").
		WithTeamId(32).
		WithTeamName("nameteam1").
		WithQ1time(time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local)).
		WithQ2time(time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local)).
		WithQ3time(time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local)).
		WithGPId(11).
		WithGPName("gpName").
		Build()

	s.qRepoMock.On("Create", &qual).Return(1, nil)
	s.qRepoMock.On("GetQualResultById", qualView.ID).Return(&qualView, nil)

	id, err := s.uc.Create(&qual)
	t.Assert().NoError(err)

	res, err := s.uc.GetQualResultById(id)
	t.Assert().NoError(err)

	t.Assert().Equal(1, id)
	t.Assert().Equal(&qualView, res)
}

func (s *QualResultTestSuite) TestUpdate(t provider.T) {
	qual, _ := s.qBuilder.
		WithID(1).
		WithDriverPlace(1).
		WithDriverId(12).
		WithDriverName("name1").
		WithTeamId(32).
		WithTeamName("nameteam1").
		WithQ1time(time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local)).
		WithQ2time(time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local)).
		WithQ3time(time.Date(2023, time.Month(2), 2, 0, 1, 12, 123, time.Local)).
		WithGPId(11).
		WithGPName("gpName").
		Build()

	s.qRepoMock.On("Update", &qual).Return(nil)

	err := s.uc.Update(1, &qual)
	t.Assert().NoError(err)
}

func (s *QualResultTestSuite) TestGetQualResultByIdWithError(t provider.T) {
	id := 1

	s.qRepoMock.On("GetQualResultById", id).Return(nil, fmt.Errorf("unexpected error"))
	res, err := s.uc.GetQualResultById(id)

	t.Assert().Error(err)
	t.Assert().Nil(res)
}
