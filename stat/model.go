package main

import "time"

type Stat struct {
	Id			int			`json:"id"`
	RoundId		int			`json:"round_id"`
	Player		Player		`json:"player"`
	Game		Game		`json:"game"`
	Result		bool		`json:"result"`
	Date		time.Time 	`json:"date"`
}

type Stats []Stat
