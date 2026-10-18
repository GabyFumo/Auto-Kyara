package models

import "time"

type Discord struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type Trust struct {
	Enforced bool  `json:"enforced"`
	MFA      bool  `json:"mfa"`
	Ranking  bool  `json:"ranking"`
	Referral bool  `json:"referral"`
	Market   bool  `json:"market"`
	Limited  any   `json:"limited"`
	AltSide  any   `json:"altSide"`
	Tips     []any `json:"tips"`
	TipCodes []any `json:"tipCodes"`
}

type User struct {
	ID                string         `json:"id"`
	Pseudo            string         `json:"pseudo"`
	Discord           Discord        `json:"discord"`
	Chakra            int            `json:"chakra"`
	ChakraHeld        int            `json:"chakraHeld"`
	FreePacks         int            `json:"freePacks"`
	FreePacksMax      int            `json:"freePacksMax"`
	FreeBoost         any            `json:"freeBoost"`
	Membership        any            `json:"membership"`
	NextFreeAt        time.Time      `json:"nextFreeAt"`
	Pity              int            `json:"pity"`
	BoostersOpened    int            `json:"boostersOpened"`
	BonusPacks        int            `json:"bonusPacks"`
	GoldenIn          int            `json:"goldenIn"`
	Grants            map[string]int `json:"grants"`
	CreatedAt         time.Time      `json:"createdAt"`
	IsAdmin           bool           `json:"isAdmin"`
	IsChatModerator   bool           `json:"isChatModerator"`
	IsAppealModerator bool           `json:"isAppealModerator"`
	Ban               any            `json:"ban"`
	DeletionAt        any            `json:"deletionAt"`
	Frame             any            `json:"frame"`
	Trust             Trust          `json:"trust"`
}
