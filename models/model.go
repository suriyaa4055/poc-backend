package models
//Lists the Head unit Generation and its version
type TestBenchTable struct {
	HuGen     string
	HuVersion string
}
//All Details Of Head unit are included below to be displayed on Home Page.
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
