package models

type PowerRepository struct {
	ActivePower int `json:"active_power" bson:"active_power" db:"active_power"`
	PowerInput  int `json:"power_input" bson:"power_input"  db:"power_input"`
}
