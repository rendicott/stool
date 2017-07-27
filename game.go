package main

type Game struct {
	Id			int			`json:"id"`
	Name 		string		`json:"name"`
}

type Games []Game