package models

import "time"

type PowerPlant struct {
	ID   int
	Name string
}

type Measurement struct {
	ID          int
	ProjectID   int
	EnergyInKwh int
	DateTime    time.Time
}
