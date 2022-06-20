package main

type Car struct {
	ID          string `json:"id"`
	Name        string `json:"string"`
	Colour      string `json:"colour"`
	Year        string `json:"year"`
	Engine      string `json:"engine"`
	Doors       int    `json:"doors"`
	Warranty    int    `json:"warranty"`
	AutoPilot   bool   `json:"autoPilot"`
	AutoParking bool   `json:"autoParking"`
}

var carList = []Car{
	{
		ID:          "001",
		Name:        "Model 2",
		Colour:      "RED",
		Year:        "2022",
		Engine:      "Hybrid",
		Doors:       2,
		Warranty:    5,
		AutoPilot:   true,
		AutoParking: true,
	},
	{
		ID:          "002",
		Name:        "Model 3",
		Colour:      "BLUE",
		Year:        "2022",
		Engine:      "Hybrid",
		Doors:       4,
		Warranty:    5,
		AutoPilot:   true,
		AutoParking: true,
	},
}
