package models

type MerchantCampaignResponse struct {
	MerchantName     string `json:"merchant_name"`
	CampaignName     string `json:"campaign_name"`
	DSC              string `json:"dsc"`
	CampaignImageURL string `json:"campaign_image_url"`
	ProviderName     string `json:"provider_name"`
	RequiredObNum    int64  `json:"required_ob_num"`
	RewardScore      int64  `json:"reward_score"`
	ItemName         string `json:"item_name"`
}

//func (m *MerchantCampaignResponse) ToUserResponse(providers []*models.Provider,
//	campaigns []*models.Campaign,
//	merchants []*models.Merchant,
//	items []*models.Item,
//	rules []*models.Rule,
//	merchantCampaign []*models.MerchantCampaign) ([]*MerchantCampaignResponse, error) {
//
//	var res []*MerchantCampaignResponse
//
//	for _, mcampaigns := range merchantCampaign {
//
//	}
//
//}
