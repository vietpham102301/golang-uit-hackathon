package repos

import "gorm.io/gorm"

type Repo struct {
	db *gorm.DB
}

func NewSQLRepo(db *gorm.DB) IRepo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Book() IBookRepo {
	return NewBookSQLRepo(r.db)
}
