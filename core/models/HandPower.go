package models

type PowerHandler struct {
	ActivePower int `json:"active_power"`
	PowerInput  int `json:"power_input"`
}
type PowerGetSumHandler struct {
	Sum int `json:"sum"`
}
