package models

type PowerSrv struct {
	ActivePower int `json:"active_power"`
	PowerInput  int `json:"power_input"`
}

type PowerGetSumSrv struct {
	Sum int `json:"sum"`
}
