package crud

import "gorm.io/gorm"

type CrudRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) CrudRepository {
	return CrudRepository{
		db: db,
	}
}
