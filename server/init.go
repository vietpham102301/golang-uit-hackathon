package server

import (
	hdl "github.com/vietpham1023/golang-uit-hackathon/handler"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
	"github.com/vietpham1023/golang-uit-hackathon/internal/services/book"
	campaign2 "github.com/vietpham1023/golang-uit-hackathon/internal/services/campaign"

	item2 "github.com/vietpham1023/golang-uit-hackathon/internal/services/item"
	merchant2 "github.com/vietpham1023/golang-uit-hackathon/internal/services/merchant"

	"github.com/vietpham1023/golang-uit-hackathon/internal/services/merchant_campaign"
	provider2 "github.com/vietpham1023/golang-uit-hackathon/internal/services/provider"
	rule2 "github.com/vietpham1023/golang-uit-hackathon/internal/services/rule"
	user2 "github.com/vietpham1023/golang-uit-hackathon/internal/services/user"
)

func (s *Server) initServices(repo repos.IRepo) *ServiceList {
	book := book.NewBook(s.cfg, repo)
	user := user2.NewUser(s.cfg, repo)
	merchantCampaign := merchant_campaign.NewMerchantCampaign(s.cfg, repo)
	campaign := campaign2.NewCampaign(s.cfg, repo)
	rule := rule2.NewRule(s.cfg, repo)
	item := item2.NewItem(s.cfg, repo)
	merchant := merchant2.NewMerchant(s.cfg, repo)
	provider := provider2.NewProvider(s.cfg, repo)
	return &ServiceList{
		book:             book,
		user:             user,
		merchantCampaign: merchantCampaign,
		campaign:         campaign,
		rule:             rule,
		item:             item,
		merchant:         merchant,
		provider:         provider,
	}
}

func (s *Server) initRouters(serviceList *ServiceList) {

	handler := hdl.NewHandler(
		s.cfg,
		s.router,
		serviceList.book,
		serviceList.user,
		serviceList.merchantCampaign,
		serviceList.campaign,
		serviceList.rule,
		serviceList.item,
		serviceList.merchant,
		serviceList.provider,
	)

	handler.ConfigureAPIRoute(s.router)
	handler.ConfigureAPIMerchantRoute(s.router)

}
