package models

type Featured struct {
	GpID   int64 `json:"gp_id" db:"gp_id"`
	UserID int64 `json:"user_id" db:"user_id"`
}

type FeaturedUsecaseI interface {
	GetAll(user_id int64) ([]*Featured, error)
	Create(featured *Featured) (int, error)
	Delete(featured *Featured) error
}

type FeaturedRepositoryI interface {
	GetAll(user_id int64) ([]*Featured, error)
	Create(featured *Featured) (int, error)
	Delete(featured *Featured) error
}
