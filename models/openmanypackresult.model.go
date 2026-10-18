package models

type OpenManyPacks struct {
	Results []OpenOnePack `json:"results"`
	Stopped any           `json:"stopped"`
}
