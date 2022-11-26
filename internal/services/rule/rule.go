package rule

import (
	"fmt"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
)

type IRule interface {
	ListRuleByIDs(IDs []int64) ([]*models.Rule, error)
	CreateRule(rule *models.Rule) (*models.Rule, error)
}

type Rule struct {
	cfg  *config.AppConfig
	repo repos.IRepo
}

func NewRule(appConfig *config.AppConfig, repo repos.IRepo) IRule {
	return &Rule{
		cfg:  appConfig,
		repo: repo,
	}
}

func (r Rule) ListRuleByIDs(IDs []int64) ([]*models.Rule, error) {
	records, err := r.repo.Rule().ListRuleByIDs(IDs)
	if err != nil {
		fmt.Printf("err :%v\n", err)
		return nil, err
	}
	return records, nil
}

func (r Rule) CreateRule(rule *models.Rule) (*models.Rule, error) {
	record, err := r.repo.Rule().Create(rule)
	return record, err
}
