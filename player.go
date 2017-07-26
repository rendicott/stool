package main

type Player struct {
	Id		int		`json:"id"`
	Name	string	`json:"name"`
}

type Players []Player