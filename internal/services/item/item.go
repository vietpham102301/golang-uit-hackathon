package item

import (
	"fmt"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
)

type IItem interface {
	ListItemByIDs(IDs []int64) ([]*models.Item, error)
	GetItemsByProviderID(providerID int) ([]*models.Item, error)
}

type Item struct {
	cfg  *config.AppConfig
	repo repos.IRepo
}

func NewItem(appConfig *config.AppConfig, repo repos.IRepo) IItem {
	return &Item{
		cfg:  appConfig,
		repo: repo,
	}
}

func (r Item) ListItemByIDs(IDs []int64) ([]*models.Item, error) {
	records, err := r.repo.Item().ListItemByIDs(IDs)
	if err != nil {
		fmt.Printf("err :%v\n", err)
		return nil, err
	}
	return records, nil
}

func (r Item) GetItemsByProviderID(providerID int) ([]*models.Item, error) {
	records, err := r.repo.Item().GetItemByProviderID(providerID)
	if err != nil {
		fmt.Printf("err :%v\n", err)
		return nil, err
	}
	return records, nil
}
