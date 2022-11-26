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

func (r *Repo) User() IUserRepo {
	return NewUserSQLRepo(r.db)
}

func (r *Repo) Citizen() ICitizenRepo {
	return NewCitizenSQLRepo(r.db)
}

func (r *Repo) MerchantCampaign() IMerchantCampaignRepo {
	return NewMerchantCampaignSQLRepo(r.db)
}

func (r *Repo) Campaign() ICampaignRepo {
	return NewCampaignSQLRepo(r.db)
}

func (r *Repo) Rule() IRuleRepo {
	return NewRuleSQLRepo(r.db)
}

func (r *Repo) Merchant() IMerchantRepo {
	return NewMerchantSQLRepo(r.db)
}
