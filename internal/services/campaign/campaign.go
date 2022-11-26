package campaign

import (
	"fmt"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
)

type ICampaign interface {
	ListCampaignByIDs(IDs []int64) ([]*models.Campaign, error)
}

type Campaign struct {
	cfg  *config.AppConfig
	repo repos.IRepo
}

func NewCampaign(appConfig *config.AppConfig, repo repos.IRepo) ICampaign {
	return &Campaign{
		cfg:  appConfig,
		repo: repo,
	}
}

func (c *Campaign) ListCampaignByIDs(IDs []int64) ([]*models.Campaign, error) {
	records, err := c.repo.Campaign().ListCampaignByIDs(IDs)
	if err != nil {
		fmt.Printf("fail to get campaign list from ids err: %v\n", err)
		return nil, err
	}
	return records, nil
}
