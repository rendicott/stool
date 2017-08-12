package main

type Outcome struct {
	Id      int     `json:"id"`
	Game    Game    `json:"game"`
	Player  Player  `json:"player"`
	Win     bool    `json:"win"`
}

type Outcomes []Outcomes