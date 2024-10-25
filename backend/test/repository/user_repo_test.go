package test

import (
	"testing"

	"project/internal/pkg/models"
	r "project/internal/pkg/user/repository/postgresql"
	"project/test/testutils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

func MapUser(user *models.User) *sqlmock.Rows {
	return sqlmock.NewRows([]string{"user_id", "login", "password", "role"}).
		AddRow(user.ID, user.Login, user.Password, user.Role)
}

type UserRepoTestSuite struct {
	suite.Suite
	uBuilder *testutils.UserBuilder
}

func TestUserTestSuite(t *testing.T) {
	suite.RunSuite(t, new(UserRepoTestSuite))
}

func (s *UserRepoTestSuite) BeforeEach(t provider.T) {
	s.uBuilder = testutils.NewUserBuilder()
}

func (s *UserRepoTestSuite) TestGetUserById(t provider.T) {
	id := 1
	user := s.uBuilder.WithID(id).WithLogin("login1").WithPassword("password1").WithRole("role").Build()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("select (.+) from users where user_id = (.+)$").
		WithArgs(id).
		WillReturnRows(MapUser(&user))

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := r.NewPsqlUserRepository(sqlxDB)

	res, err := repo.GetUserById(id)
	t.Assert().NoError(err)
	t.Assert().Equal(user, *res)
}

func (s *UserRepoTestSuite) TestGetUserByLogin(t provider.T) {
	id := 1
	login := "login"
	user := s.uBuilder.WithID(id).WithLogin(login).WithPassword("password1").WithRole("role").Build()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("select (.+) from users where login = (.+)$").
		WithArgs(login).
		WillReturnRows(MapUser(&user))

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := r.NewPsqlUserRepository(sqlxDB)

	res, err := repo.GetUserByLogin(login)
	t.Assert().NoError(err)
	t.Assert().Equal(user, *res)
}

func (s *UserRepoTestSuite) TestCreate(t provider.T) {
	id := 1
	login := "login"
	user := s.uBuilder.WithID(id).WithLogin(login).WithPassword("password1").WithRole("role").Build()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("insert into users").
		WithArgs(user.Login, user.Password, user.Role).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(user.ID))

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := r.NewPsqlUserRepository(sqlxDB)

	res, err := repo.Create(&user)
	t.Assert().NoError(err)
	t.Assert().Equal(id, res)
}

func (s *UserRepoTestSuite) TestUpdate(t provider.T) {
	id := 1
	login := "login"
	user := s.uBuilder.WithID(id).WithLogin(login).WithPassword("password1").WithRole("role").Build()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("update users").
		WithArgs(user.Login, user.Password, user.Role, user.ID).RowsWillBeClosed()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := r.NewPsqlUserRepository(sqlxDB)

	err = repo.Update(&user)
	t.Assert().NoError(err)
}

func (s *UserRepoTestSuite) TestDelete(t provider.T) {
	id := 1
	login := "login"
	user := s.uBuilder.WithID(id).WithLogin(login).WithPassword("password1").WithRole("role").Build()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("delete from users").
		WithArgs(user.ID)

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := r.NewPsqlUserRepository(sqlxDB)

	err = repo.Delete(user.ID)
	t.Assert().NoError(err)
}

// TODO: достаточно даже одного класса для тестирования(например юзер :))
