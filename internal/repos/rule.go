package repos

import (
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"gorm.io/gorm"
)

type RuleSQLRepo struct {
	db *gorm.DB
}

func NewRuleSQLRepo(db *gorm.DB) IRuleRepo {
	return &RuleSQLRepo{
		db: db,
	}
}

func (r RuleSQLRepo) ListRuleByIDs(IDs []int64) ([]*models.Rule, error) {
	var records []*models.Rule
	err := r.db.Where("id in (?)", IDs).First(&records).Error
	return records, err
}
