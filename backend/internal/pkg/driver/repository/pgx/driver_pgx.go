package pgx

import (
	"context"
	"project/internal/pkg/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type psqlDriverRepository struct {
	db *pgxpool.Pool
}

func NewPsqlDriverRepositoryPGX(db *pgxpool.Pool) models.DriverRepositoryI {
	return &psqlDriverRepository{
		db: db,
	}
}

func (pgRepo *psqlDriverRepository) GetAll() ([]*models.Driver, error) {
	drivers := []*models.Driver{}
	rows, err := pgRepo.db.Query(context.Background(),
		"select driver_id, driver_name, driver_country, driver_birth_date "+
			"from drivers")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		driver := &models.Driver{}
		err = rows.Scan(&driver)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (pgRepo *psqlDriverRepository) GetDriverById(id int) (*models.Driver, error) {
	driver := &models.Driver{}
	row := pgRepo.db.QueryRow(
		context.Background(),
		"select driver_id, driver_name, driver_country, driver_birth_date "+
			"from drivers "+
			"where driver_id = $1",
		id)
	err := row.Scan(&driver)
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (pgRepo *psqlDriverRepository) GetDriversOfSeason(season int) ([]*models.Driver, error) {
	drivers := []*models.Driver{}

	return drivers, nil
}

func (pgRepo *psqlDriverRepository) GetDriversStanding() ([]*models.Standings, error) {
	standings := []*models.Standings{}

	return standings, nil
}

func (pgRepo *psqlDriverRepository) Create(driver *models.Driver) (int, error) {
	var id int
	_, err := pgRepo.db.Exec(context.Background(),
		"insert into drivers (driver_name, driver_country, driver_birth_date) "+
			"values ($1, $2, $3) "+
			"returning driver_id",
		driver.Name,
		driver.Country,
		driver.BirthDate,
	)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (pgRepo *psqlDriverRepository) Update(newDriver *models.Driver) error {
	_, err := pgRepo.db.Exec(context.Background(),
		"update drivers "+
			"set driver_name = $1, "+
			"driver_country = $2, "+
			"driver_birth_date = $3 "+
			"where driver_id = $4",
		newDriver.Name,
		newDriver.Country,
		newDriver.BirthDate,
		newDriver.ID)
	if err != nil {
		return err
	}
	return nil
}

func (pgRepo *psqlDriverRepository) Delete(id int) error {
	_, err := pgRepo.db.Exec(context.Background(),
		"delete from drivers "+
			"where driver_id = $1",
		id)
	if err != nil {
		return err
	}
	return nil
}

func (pgRepo *psqlDriverRepository) LinkDriverTeam(new *models.DriversTeams) error {
	_, err := pgRepo.db.Exec(context.Background(),
		"insert into teamsdrivers (driver_id, team_id, team_driver_season) "+
			"values ($1, $2, $3)",
		new.DriverId,
		new.TeamId,
		new.Season,
	)
	if err != nil {
		return err
	}
	return nil
}
