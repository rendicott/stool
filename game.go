package main

type Game struct {
	Id			int			`json:"id"`
	Name 		string		`json:"name"`
	Players		int			`json:"players"`
}

type Games []Game