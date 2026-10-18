package models

type OpenOnePack struct {
	Cards        []Card `json:"cards"`
	User         User   `json:"user"`
	PackID       string `json:"packId"`
	ShareID      string `json:"shareId"`
	Golden       bool   `json:"golden"`
	GodPack      bool   `json:"godPack"`
	Pity         bool   `json:"pity"`
	FeaturedPity bool   `json:"featuredPity"`
	PaidWith     string `json:"paidWith"`
	Theme        any    `json:"theme"`
	Unlocked     []any  `json:"unlocked"`
	Claimable    int    `json:"claimable"`
}
