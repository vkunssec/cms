package entity

type ResponseIndex struct {
	Uptime      string `json:"uptime"`
	Message     string `json:"message"`
	Version     string `json:"version"`
	Environment string `json:"env"`
	Date        string `json:"date"`
}
