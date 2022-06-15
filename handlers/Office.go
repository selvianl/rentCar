package handlers

type Office struct {
	Id      uint    `json:"id"`
	Vendor Vendor `json:"vendor"`
	Location Location `json:"product"`
}