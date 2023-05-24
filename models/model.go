package models

type TestBenchTable struct {
	HuVersion string `json:"huVersion"`
	HuGen     string `json:"huGen"`
}

type Gen20x struct {
	Version   []string  `json:"version"`
	SwVersion []float64 `json:"swVersion"`
}

type NTG struct {
	Version   []string  `json:"version"`
	SwVersion []float64 `json:"swVersion"`
}

type HU struct {
	Gen20x *Gen20x `json:"gen20x"`
	NTG    *NTG    `json:"ntg"`
}
