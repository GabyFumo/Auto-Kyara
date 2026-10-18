package models

type Card struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Native     string `json:"native"`
	Image      string `json:"image"`
	Favourites int    `json:"favourites"`
	Rank       int    `json:"rank"`
	Rarity     string `json:"rarity"`
	Num        string `json:"num"`
	Series     string `json:"series"`
	Pow        int    `json:"pow"`
	Aura       int    `json:"aura"`
	IsNew      bool   `json:"isNew"`
	Featured   bool   `json:"featured"`
}
