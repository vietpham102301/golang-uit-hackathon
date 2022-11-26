package models

type Rule struct {
	ID            int64 `json:"id"`
	ItemId        int64 `json:"item_id"`
	RequiredObNum int64 `json:"required_ob_num"`
	RewardScore   int64 `json:"reward_score"`
}
