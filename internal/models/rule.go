package models

type Rule struct {
	ID                int64 `json:"id"`
	ItemId            int64 `json:"item_id"`
	RequiredObjectNum int64 `json:"required_object_num"`
	RewardScore       int64 `json:"reward_score"`
}
