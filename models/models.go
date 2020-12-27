package models

import "time"

type Team struct {
	Name             string
	Location         string
	Rank             int
	FullRecord       string
	HomeRecord       string
	AwayRecord       string
	ConferenceRecord string
	Score            string
}

type Game struct {
	Home   Team
	Away   Team
	Time   time.Time
	Status string
}

type Standing struct {
	Team             Team
	Position         int
	FullRecord       string
	HomeRecord       string
	AwayRecord       string
	ConferenceRecord string
	GamesBehind      int
}
