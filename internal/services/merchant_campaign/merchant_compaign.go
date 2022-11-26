package merchant_campaign

import (
	"context"
	"fmt"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	handlerModels "github.com/vietpham1023/golang-uit-hackathon/handler/models"
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
)

type IMerchantCampaign interface {
	Create(ctx context.Context, record *handlerModels.MerchantCampaign) error
	ListByFilter(data map[string]string) ([]*models.MerchantCampaignResponse, error)
}

type MerchantCampaign struct {
	cfg  *config.AppConfig
	repo repos.IRepo
}

func NewMerchantCampaign(appConfig *config.AppConfig, repo repos.IRepo) IMerchantCampaign {
	return &MerchantCampaign{
		cfg:  appConfig,
		repo: repo,
	}
}

func (m *MerchantCampaign) ListByFilter(data map[string]string) ([]*models.MerchantCampaignResponse, error) {
	page, size, filter := getFilterList(data)
	records, err := m.repo.MerchantCampaign().ListMerchantCampaignRepo(page, size, filter)

	if err != nil {
		fmt.Printf("list book fail with err: %v\n", err)
		return nil, err
	}
	return records, nil
}

func (m *MerchantCampaign) Create(ctx context.Context, record *handlerModels.MerchantCampaign) error {
	err := m.repo.MerchantCampaign().Create(ctx, record)
	if err != nil {
		return err
	}

	return nil
}
