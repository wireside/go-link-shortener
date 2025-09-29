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

func (repo LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.database.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repo LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.database.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}
