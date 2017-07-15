package main

type Game struct {
	Name 		string		`json:"name"`
	Players		int			`json:"players"`
	Id 			int			`json:"id"`
}

type Games []Game