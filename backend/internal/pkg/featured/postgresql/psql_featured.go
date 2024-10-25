package postgresql

import (
	"fmt"
	"project/internal/pkg/models"

	"github.com/jmoiron/sqlx"
)

type psqlGPRepository struct {
	db *sqlx.DB
}

func NewPsqlGPRepository(db *sqlx.DB) models.FeaturedRepositoryI {
	return &psqlGPRepository{
		db: db,
	}
}

func (pgRepo *psqlGPRepository) GetAll(user_id int64) ([]*models.Featured, error) {
	featured := []*models.Featured{}
	rows, err := pgRepo.db.Queryx(
		"select gp_id, user_id "+
			"from gp_users "+
			"where user_id = $1", user_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		feature_temp := &models.Featured{}
		err = rows.StructScan(feature_temp)
		if err != nil {
			return nil, err
		}
		featured = append(featured, feature_temp)
	}
	return featured, nil
}

func (pgRepo *psqlGPRepository) Create(featured *models.Featured) (int, error) {
	var id int
	err := pgRepo.db.QueryRow(
		"insert into gp_users (gp_id, user_id)"+
			"values ($1, $2) "+
			"returning gp_id",
		featured.GpID,
		featured.UserID,
	).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}

func (pgRepo *psqlGPRepository) Delete(feature *models.Featured) error {
	_, err := pgRepo.db.Exec(
		"delete from grandprix "+
			"where gp_id = $1 and user_id = $2",
		feature.GpID, feature.UserID)
	if err != nil {
		return err
	}
	return nil
}
