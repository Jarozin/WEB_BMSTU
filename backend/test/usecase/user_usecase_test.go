package test

import (
	"project/internal/pkg/models"
	"testing"

	"project/internal/pkg/models/mocks"
	"project/internal/pkg/user/usecase"
	"project/test/testutils"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type UserTestSuite struct {
	suite.Suite
	uc        models.UserUsecaseI
	uRepoMock *mocks.UserRepositoryI
	uBuilder  *testutils.UserBuilder
}

func TestUserTestSuite(t *testing.T) {
	suite.RunSuite(t, new(UserTestSuite))
}

func (s *UserTestSuite) BeforeEach(t provider.T) {
	s.uRepoMock = mocks.NewUserRepositoryI(t)
	s.uc = usecase.NewUserUsecase(s.uRepoMock)
	s.uBuilder = testutils.NewUserBuilder()
}

func (s *UserTestSuite) TestGetUserById(t provider.T) {
	user := s.uBuilder.WithID(1).WithLogin("login1").WithPassword("password1").WithRole("role").Build()

	s.uRepoMock.On("GetUserById", user.ID).Return(&user, nil)
	res, err := s.uc.GetUserById(user.ID)

	t.Assert().NoError(err)
	t.Assert().Equal(user, *res)
}

func (s *UserTestSuite) TestGetUserByLogin(t provider.T) {
	user := s.uBuilder.WithID(1).WithLogin("login1").WithPassword("password1").WithRole("role").Build()

	s.uRepoMock.On("GetUserByLogin", user.Login).Return(&user, nil)
	res, err := s.uc.GetUserByLogin(user.Login)

	t.Assert().NoError(err)
	t.Assert().Equal(user, *res)
}

func (s *UserTestSuite) TestAuthenticate(t provider.T) {
	user := s.uBuilder.WithID(1).WithLogin("login1").WithPassword("password1").WithRole("role").Build()

	s.uRepoMock.On("GetUserByLogin", user.Login).Return(&user, nil)

	ok, err := s.uc.Authenticate(user.Login, user.Password)
	t.Assert().NoError(err)

	t.Assert().Equal(true, ok)
}

func (s *UserTestSuite) TestAuthenticateFalse(t provider.T) {
	user := s.uBuilder.WithID(1).WithLogin("login1").WithPassword("password1").WithRole("role").Build()

	s.uRepoMock.On("GetUserByLogin", user.Login).Return(&user, nil)

	ok, err := s.uc.Authenticate(user.Login, "other_password")
	t.Assert().NoError(err)

	t.Assert().Equal(false, ok)
}

func (s *UserTestSuite) TestCreate(t provider.T) {
	user := s.uBuilder.WithID(1).WithLogin("login1").WithPassword("password1").WithRole("role").Build()

	s.uRepoMock.On("Create", &user).Return(1, nil)
	s.uRepoMock.On("GetUserById", user.ID).Return(&user, nil)

	id, err := s.uc.Create(&user)
	t.Assert().NoError(err)

	res, err := s.uc.GetUserById(user.ID)
	t.Assert().NoError(err)

	t.Assert().Equal(1, id)
	t.Assert().Equal(&user, res)
}

func (s *UserTestSuite) TestUpdate(t provider.T) {
	user := s.uBuilder.WithID(1).WithLogin("login1").WithPassword("password1").WithRole("role").Build()

	s.uRepoMock.On("Update", &user).Return(nil)
	s.uRepoMock.On("GetUserById", user.ID).Return(&user, nil)

	err := s.uc.Update(1, &user)
	t.Assert().NoError(err)

	res, err := s.uc.GetUserById(user.ID)
	t.Assert().NoError(err)
	t.Assert().Equal(&user, res)
}
