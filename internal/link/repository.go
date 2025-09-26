package link

import (
	"go-adv-demo/pkg/db"
)

type LinkRepository struct {
	database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		database,
	}
}

func (repo LinkRepository) Create(link *Link) *Link {
	return nil
}
