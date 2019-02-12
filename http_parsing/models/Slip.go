package models

type Response struct {
	Slip Slip `json:"slip"`
}
type Slip struct {
	Advice string `json:"advice"`
	Slip_ID string `json:"slip_id"`
}
