package models

type TestBenchTable struct {
	HuGen     string
	HuVersion string
}

type Details struct {
	Vin       string `json:"vin"`
	Region    string `json:"region"`
	HUGen     string `json:"gen"`
	HUVersion string `json:"version"`
	HuRelease string `json:"release"`

	Stage     string `json:"stage"`
	Cubicle   string `json:"Cubicle"`
	SwVersion string `json:"SwVersion"`
	Connected bool   `json:"connected"`
}
