package repos

import (
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
)

type IRepo interface {
	Book() IBookRepo
	User() IUserRepo
	Citizen() ICitizenRepo
	MerchantCampaign() IMerchantCampaignRepo
	Campaign() ICampaignRepo
	Rule() IRuleRepo
	Item() IItemRepo
	Merchant() IMerchantRepo
	Provider() IProviderRepo
}

type IBookRepo interface {
	Create(record *models.Book) (*models.Book, error)
	GetByID(ID int) (*models.Book, error)
	ListBookByFilter(page, size int, filter map[string]interface{}) ([]*models.Book, error)
}

type IUserRepo interface {
	Create(record *models.User) (*models.User, error)
	Login(record *models.User) (*models.User, error)
	GetUserByToken(token string) (*models.User, error)
}

type ICitizenRepo interface {
	Create(record *models.Citizen) (*models.Citizen, error)
	GetByID(ID int64) (*models.Citizen, error)
}

type ICampaignRepo interface {
	ListCampaignByIDs(IDs []int64) ([]*models.Campaign, error)
	Create(record *models.Campaign) (*models.Campaign, error)
}

type IMerchantCampaignRepo interface {
	ListMerchantCampaignRepo(page int, size int, filter map[string]interface{}) ([]*models.MerchantCampaignResponse, error)
}

type IRuleRepo interface {
	ListRuleByIDs(IDs []int64) ([]*models.Rule, error)
	Create(record *models.Rule) (*models.Rule, error)
}

type IItemRepo interface {
	ListItemByIDs(IDs []int64) ([]*models.Item, error)
	GetItemByProviderID(providerID int) ([]*models.Item, error)
}

type IMerchantRepo interface {
	ListMerchantByIDs(IDs []int64) ([]*models.Merchant, error)
	Create(record *models.Merchant) (*models.Merchant, error)
	FindByID(id int) (*models.Merchant, error)
	ListByConditions(filter map[string]interface{}) ([]*models.Merchant, error)
}

type IProviderRepo interface {
	ListProviderByIDs(IDs []int64) ([]*models.Provider, error)
	Create(record *models.Provider) (*models.Provider, error)
}
