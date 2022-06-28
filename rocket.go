package main

import (
	"encoding/json"
)

type Rocket struct {
	ID             int                   `json:"id"`
	Active         bool                  `json:"active"`
	Stages         int                   `json:"stages"`
	Boosters       int                   `json:"boosters"`
	CostPerLaunch  int                   `json:"cost_per_launch"`
	SuccessRatePct int                   `json:"success_rate_pct"`
	FirstFlight    string                `json:"first_flight"`
	Country        string                `json:"country"`
	Company        string                `json:"company"`
	Height         RocketHeight          `json:"height"`
	Diameter       RocketDiameter        `json:"diameter"`
	Mass           RocketMass            `json:"mass"`
	PayloadWeights []RocketPayloadWeight `json:"payload_weights"`
	FirstStage     RocketFirstStage      `json:"first_stage"`
	SecondStage    RocketSecondStage     `json:"second_stage"`
	Engines        Engine                `json:"engines"`
	LandingLegs    LandingLegs           `json:"landing_legs"`
	Wikipedia      string                `json:"wikipedia"`
	Description    string                `json:"description"`
	RocketID       string                `json:"rocket_id"`
	RocketName     string                `json:"rocket_name"`
	RocketType     string                `json:"rocket_type"`
}

type RocketV2 struct {
	ID            int                            `json:"id"`
	RocketID      string                         `json:"rocket_id"`
	PayloadWeight map[string]RocketPayloadWeight `json:"payload_weights"`
}

type RocketHeight struct {
	Meters float64 `json:"meters"`
	Feet   float64 `json:"feet"`
}

type RocketDiameter struct {
	Meters float64 `json:"meters"`
	Feet   float64 `json:"feet"`
}

type RocketMass struct {
	Kg int `json:"kg"`
	Lb int `json:"lb"`
}

type RocketPayloadWeight struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Kg   int    `json:"kg"`
	Lb   int    `json:"lb"`
}

type RocketFirstStage struct {
	Reusable       bool    `json:"reusable"`
	Engines        int     `json:"engines"`
	FuelAmountTons float64 `json:"fuel_amount_tons"`
	BurnTimeSec    int     `json:"burn_time_sec"`
	ThrustSeaLevel Trust   `json:"thrust_sea_level"`
	ThrustVacuum   Trust   `json:"thrust_vacuum"`
}

type Trust struct {
	KN  int `json:"kN"`
	Lbf int `json:"lbf"`
}

type RocketSecondStage struct {
	Engines        int                `json:"engines"`
	FuelAmountTons float64            `json:"fuel_amount_tons"`
	BurnTimeSec    int                `json:"burn_time_sec"`
	Thrust         Trust              `json:"thrust"`
	Payloads       SecondStagePayload `json:"payloads"`
}

type SecondStagePayload struct {
	Option1          string           `json:"option_1"`
	CompositeFairing CompositeFairing `json:"composite_fairing"`
}
type CompositeFairing struct {
	Height   RocketHeight   `json:"height"`
	Diameter RocketDiameter `json:"diameter"`
}

type Engine struct {
	Number         int     `json:"number"`
	Type           string  `json:"type"`
	Version        string  `json:"version"`
	Layout         string  `json:"layout"`
	EngineLossMax  float64 `json:"engine_loss_max"`
	Propellant1    string  `json:"propellant_1"`
	Propellant2    string  `json:"propellant_2"`
	ThrustSeaLevel Trust   `json:"thrust_sea_level"`
	ThrustVacuum   Trust   `json:"thrust_vacuum"`
	ThrustToWeight float64 `json:"thrust_to_weight"`
}

type LandingLegs struct {
	Number   int         `json:"number"`
	Material interface{} `json:"material"`
}

// json taken from https://docs.spacexdata.com/#5fcdb875-914f-4aef-a932-254397cf147a
var rocketsJSON = []byte(`
[
  {
    "id": 1,
    "active": false,
    "stages": 2,
    "boosters": 0,
    "cost_per_launch": 6700000,
    "success_rate_pct": 40,
    "first_flight": "2006-03-24",
    "country": "Republic of the Marshall Islands",
    "company": "SpaceX",
    "height": {
      "meters": 22.25,
      "feet": 73
    },
    "diameter": {
      "meters": 1.68,
      "feet": 5.5
    },
    "mass": {
      "kg": 30146,
      "lb": 66460
    },
    "payload_weights": [
      {
        "id": "leo",
        "name": "Low Earth Orbit",
        "kg": 450,
        "lb": 992
      }
    ],
    "first_stage": {
      "reusable": false,
      "engines": 1,
      "fuel_amount_tons": 44.3,
      "burn_time_sec": 169,
      "thrust_sea_level": {
        "kN": 420,
        "lbf": 94000
      },
      "thrust_vacuum": {
        "kN": 480,
        "lbf": 110000
      }
    },
    "second_stage": {
      "engines": 1,
      "fuel_amount_tons": 3.38,
      "burn_time_sec": 378,
      "thrust": {
        "kN": 31,
        "lbf": 7000
      },
      "payloads": {
        "option_1": "composite fairing",
        "composite_fairing": {
          "height": {
            "meters": 3.5,
            "feet": 11.5
          },
          "diameter": {
            "meters": 1.5,
            "feet": 4.9
          }
        }
      }
    },
    "engines": {
      "number": 1,
      "type": "merlin",
      "version": "1C",
      "layout": "single",
      "engine_loss_max": 0,
      "propellant_1": "liquid oxygen",
      "propellant_2": "RP-1 kerosene",
      "thrust_sea_level": {
        "kN": 420,
        "lbf": 94000
      },
      "thrust_vacuum": {
        "kN": 480,
        "lbf": 110000
      },
      "thrust_to_weight": 96
    },
    "landing_legs": {
      "number": 0,
      "material": null
    },
    "wikipedia": "https://en.wikipedia.org/wiki/Falcon_1",
    "description": "The Falcon 1 was an expendable launch system privately developed and manufactured by SpaceX during 2006-2009. On 28 September 2008, Falcon 1 became the first privately-developed liquid-fuel launch vehicle to go into orbit around the Earth.",
    "rocket_id": "falcon1",
    "rocket_name": "Falcon 1",
    "rocket_type": "rocket"
  },
  {
    "id": 2,
    "active": true,
    "stages": 2,
    "boosters": 0,
    "cost_per_launch": 50000000,
    "success_rate_pct": 97,
    "first_flight": "2010-06-04",
    "country": "United States",
    "company": "SpaceX",
    "height": {
      "meters": 70,
      "feet": 229.6
    },
    "diameter": {
      "meters": 3.7,
      "feet": 12
    },
    "mass": {
      "kg": 549054,
      "lb": 1207920
    },
    "payload_weights": [
      {
        "id": "leo",
        "name": "Low Earth Orbit",
        "kg": 22800,
        "lb": 50265
      },
      {
        "id": "gto",
        "name": "Geosynchronous Transfer Orbit",
        "kg": 8300,
        "lb": 18300
      },
      {
        "id": "mars",
        "name": "Mars Orbit",
        "kg": 4020,
        "lb": 8860
      }
    ],
    "first_stage": {
      "reusable": true,
      "engines": 9,
      "fuel_amount_tons": 385,
      "burn_time_sec": 162,
      "thrust_sea_level": {
        "kN": 7607,
        "lbf": 1710000
      },
      "thrust_vacuum": {
        "kN": 8227,
        "lbf": 1849500
      }
    },
    "second_stage": {
      "engines": 1,
      "fuel_amount_tons": 90,
      "burn_time_sec": 397,
      "thrust": {
        "kN": 934,
        "lbf": 210000
      },
      "payloads": {
        "option_1": "dragon",
        "option_2": "composite fairing",
        "composite_fairing": {
          "height": {
            "meters": 13.1,
            "feet": 43
          },
          "diameter": {
            "meters": 5.2,
            "feet": 17.1
          }
        }
      }
    },
    "engines": {
      "number": 9,
      "type": "merlin",
      "version": "1D+",
      "layout": "octaweb",
      "engine_loss_max": 2,
      "propellant_1": "liquid oxygen",
      "propellant_2": "RP-1 kerosene",
      "thrust_sea_level": {
        "kN": 845,
        "lbf": 190000
      },
      "thrust_vacuum": {
        "kN": 914,
        "lbf": 205500
      },
      "thrust_to_weight": 180.1
    },
    "landing_legs": {
      "number": 4,
      "material": "carbon fiber"
    },
    "wikipedia": "https://en.wikipedia.org/wiki/Falcon_9",
    "description": "Falcon 9 is a two-stage rocket designed and manufactured by SpaceX for the reliable and safe transport of satellites and the Dragon spacecraft into orbit.",
    "rocket_id": "falcon9",
    "rocket_name": "Falcon 9",
    "rocket_type": "rocket"
  },
  {
    "id": 3,
    "active": true,
    "stages": 2,
    "boosters": 2,
    "cost_per_launch": 90000000,
    "success_rate_pct": 100,
    "first_flight": "2018-02-06",
    "country": "United States",
    "company": "SpaceX",
    "height": {
      "meters": 70,
      "feet": 229.6
    },
    "diameter": {
      "meters": 12.2,
      "feet": 39.9
    },
    "mass": {
      "kg": 1420788,
      "lb": 3125735
    },
    "payload_weights": [
      {
        "id": "leo",
        "name": "Low Earth Orbit",
        "kg": 63800,
        "lb": 140660
      },
      {
        "id": "gto",
        "name": "Geosynchronous Transfer Orbit",
        "kg": 26700,
        "lb": 58860
      },
      {
        "id": "mars",
        "name": "Mars Orbit",
        "kg": 16800,
        "lb": 37040
      },
      {
        "id": "pluto",
        "name": "Pluto Orbit",
        "kg": 3500,
        "lb": 7720
      }
    ],
    "first_stage": {
      "reusable": true,
      "engines": 27,
      "fuel_amount_tons": 1155,
      "cores": 3,
      "burn_time_sec": 162,
      "thrust_sea_level": {
        "kN": 22819,
        "lbf": 5130000
      },
      "thrust_vacuum": {
        "kN": 24681,
        "lbf": 5548500
      }
    },
    "second_stage": {
      "engines": 1,
      "burn_time_sec": 397,
      "thrust": {
        "kN": 934,
        "lbf": 210000
      },
      "payloads": {
        "option_1": "dragon",
        "option_2": "composite fairing",
        "composite_fairing": {
          "height": {
            "meters": 13.1,
            "feet": 43
          },
          "diameter": {
            "meters": 5.2,
            "feet": 17.1
          }
        }
      }
    },
    "engines": {
      "number": 27,
      "type": "merlin",
      "version": "1D+",
      "layout": "octaweb",
      "engine_loss_max": 6,
      "propellant_1": "liquid oxygen",
      "propellant_2": "RP-1 kerosene",
      "thrust_sea_level": {
        "kN": 845,
        "lbf": 190000
      },
      "thrust_vacuum": {
        "kN": 914,
        "lbf": 205500
      },
      "thrust_to_weight": 180.1
    },
    "landing_legs": {
      "number": 12,
      "material": "carbon fiber"
    },
    "wikipedia": "https://en.wikipedia.org/wiki/Falcon_Heavy",
    "description": "With the ability to lift into orbit over 54 metric tons (119,000 lb)--a mass equivalent to a 737 jetliner loaded with passengers, crew, luggage and fuel--Falcon Heavy can lift more than twice the payload of the next closest operational vehicle, the Delta IV Heavy, at one-third the cost.",
    "rocket_id": "falconheavy",
    "rocket_name": "Falcon Heavy",
    "rocket_type": "rocket"
  },
  {
    "id": 4,
    "active": false,
    "stages": 2,
    "boosters": 0,
    "cost_per_launch": 7000000,
    "success_rate_pct": 0,
    "first_flight": "2019-12-01",
    "country": "United States",
    "company": "SpaceX",
    "height": {
      "meters": 106,
      "feet": 348
    },
    "diameter": {
      "meters": 9,
      "feet": 30
    },
    "mass": {
      "kg": 4400000,
      "lb": 9700000
    },
    "payload_weights": [
      {
        "id": "leo",
        "name": "Low Earth Orbit",
        "kg": 150000,
        "lb": 330000
      },
      {
        "id": "mars",
        "name": "Mars Orbit",
        "kg": 150000,
        "lb": 330000
      }
    ],
    "first_stage": {
      "reusable": true,
      "engines": 31,
      "fuel_amount_tons": 6700,
      "burn_time_sec": 0,
      "thrust_sea_level": {
        "kN": 128000,
        "lbf": 28775544
      },
      "thrust_vacuum": {
        "kN": 138000,
        "lbf": 31023634
      }
    },
    "second_stage": {
      "engines": 6,
      "fuel_amount_tons": 1100,
      "burn_time_sec": 0,
      "thrust": {
        "kN": 1900,
        "lbf": 427136
      },
      "payloads": {
        "option_1": "Spaceship",
        "option_2": "composite fairing",
        "composite_fairing": {
          "height": {
            "meters": null,
            "feet": null
          },
          "diameter": {
            "meters": null,
            "feet": null
          }
        }
      }
    },
    "engines": {
      "number": 31,
      "type": "raptor",
      "version": "",
      "layout": null,
      "engine_loss_max": null,
      "propellant_1": "liquid oxygen",
      "propellant_2": "liquid methane",
      "thrust_sea_level": {
        "kN": 1700,
        "lbf": 382175
      },
      "thrust_vacuum": {
        "kN": 1900,
        "lbf": 427136
      },
      "thrust_to_weight": null
    },
    "landing_legs": {
      "number": 4,
      "material": "carbon fiber"
    },
    "wikipedia": "https://en.wikipedia.org/wiki/BFR_(rocket)",
    "description": "BFR is a privately funded next-generation reusable launch vehicle and spacecraft system developed by SpaceX. It was announced by Elon Musk in September 2017; the first spacecraft prototype was being manufactured as of March 2018 and will begin testing in early 2019. The overall space vehicle architecture includes both launch vehicles and spacecraft that are intended to completely replace all of SpaceX's existing space hardware by the early 2020s as well as ground infrastructure for rapid launch and relaunch, and zero-gravity propellant transfer technology to be deployed in low Earth orbit (LEO). The large payload to Earth orbit of up to 150,000 kg (330,000 lb) makes BFR a super heavy-lift launch vehicle.",
    "rocket_id": "bfr",
    "rocket_name": "Big Falcon Rocket",
    "rocket_type": "rocket"
  }
]
`)

func getRocketsStruct() ([]Rocket, error) {
	var rockets []Rocket
	err := json.Unmarshal(rocketsJSON, &rockets)
	if err != nil {
		return nil, err
	}
	return rockets, nil
}

func findPayloadWeight(id string) []string {
	var results []string
	rockets, _ := getRocketsStruct()
	for _, rocket := range rockets {
		for _, payload := range rocket.PayloadWeights {
			if payload.ID == id {
				results = append(results, rocket.RocketID)
			}
		}
	}
	return results
}

func convertPayloadToMap(rockets []Rocket) []RocketV2 {
	var rocketsV2 []RocketV2
	for _, rocket := range rockets {
		rv2 := RocketV2{
			ID:            rocket.ID,
			RocketID:      rocket.RocketID,
			PayloadWeight: make(map[string]RocketPayloadWeight),
		}
		for _, payload := range rocket.PayloadWeights {
			rv2.PayloadWeight[payload.ID] = payload
		}
		rocketsV2 = append(rocketsV2, rv2)
	}

	return rocketsV2
}

func findPayloadWeightMap(id string) []string {
	var results []string
	rockets, _ := getRocketsStruct()
	rocketsV2 := convertPayloadToMap(rockets)
	for _, rocket := range rocketsV2 {
		if _, ok := rocket.PayloadWeight[id]; ok {
			results = append(results, rocket.RocketID)
		}
	}
	return results
}
