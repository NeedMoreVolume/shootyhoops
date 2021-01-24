package models

import (
	"fmt"
	"shootyhoops/models/espn"
	"strings"
	"time"
)

func competitorsToTeam(competitors []espn.Competitor) map[string]Team {
	out := make(map[string]Team, 0)

	for _, competitor := range competitors {
		var fullRecord string
		var homeRecord string
		var awayRecord string
		var conferenceRecord string
		for _, record := range competitor.Records {
			switch strings.ToLower(record.Type) {
			case "total":
				fullRecord = record.Summary
			case "home":
				homeRecord = record.Summary
			case "road":
				awayRecord = record.Summary
			case "vsconf":
				conferenceRecord = record.Summary
			}
		}
		thisTeam := Team{
			Name:             competitor.Team.Name,
			Location:         competitor.Team.Location,
			Rank:             competitor.CuratedRank.Current,
			FullRecord:       fullRecord,
			HomeRecord:       homeRecord,
			AwayRecord:       awayRecord,
			ConferenceRecord: conferenceRecord,
			Score:            competitor.Score,
		}
		out[strings.ToLower(competitor.HomeAway)] = thisTeam
	}

	return out
}

func eventToGame(event espn.Event) Game {
	teams := competitorsToTeam(event.Competitions[0].Competitors)
	homeTeam := teams["home"]
	awayTeam := teams["away"]
	eventTime, _ := time.Parse("2006-01-02T15:04Z", event.Date)

	var status string
	switch event.Status.Type.Name {
	case "STATUS_SCHEDULED":
		status = "Scheduled"
	case "STATUS_POSTPONED":
		status = "Postponed"
	case "STATUS_CANCELED":
		status = "Cancelled"
	case "STATUS_FINAL":
		status = "Final"
	case "STATUS_IN_PROGRESS":
		status = "In Progress"
	case "STATUS_HALFTIME":
		status = "Halftime"
	case "STATUS_END_PERIOD":
		status = "Period Ended"
	default:
		status = event.Status.Type.Name
	}

	return Game{
		Home:   homeTeam,
		Away:   awayTeam,
		Time:   eventTime,
		Status: status,
	}
}

func EspnGamesToGames(response espn.GameResponse) []Game {
	out := make([]Game, 0)

	for _, event := range response.Events {
		// build event into Game model
		out = append(out, eventToGame(event))
	}

	return out
}

func GamesToMessage(games []Game, withRank bool, detailed bool, scores bool, team *string) string {
	var out string

	if team != nil {
		// check if this games set has this team.
		games = FilterTeams(games, *team)

		if len(games) < 1 {
			return fmt.Sprintf("Sorry, I can't find any games for the %s on the provided date.", strings.Title(strings.ToLower(*team)))
		}
	}

	if len(games) < 1 {
		return "Sorry, I can't find any games for the provided date."
	}

	for _, game := range games {
		out += "```"
		out += FormatGameTitle(game.Home, game.Away, withRank)
		if detailed {
			out += FormatTeamDetail(game.Home)
			out += FormatTeamDetail(game.Away)
		}
		if scores {
			out += FormatGameScore(game.Home, game.Away)
		}
		out += FormatStartTime(game.Time.Local().Format(time.Kitchen))
		out += FormatStatus(game.Status)
		out += "```\n"
	}

	return out
}

func FilterTeams(games []Game, teamName string) []Game {
	var filteredGames []Game
	for _, game := range games {
		// if teamName is home team, add game and break
		if strings.Contains(game.Home.Name, strings.Title(strings.ToLower(teamName))) {
			filteredGames = append(filteredGames, game)
			break
		}

		// if teamName is away team, add game and break
		if strings.Contains(game.Away.Name, strings.Title(strings.ToLower(teamName))) {
			filteredGames = append(filteredGames, game)
			break
		}
	}

	return filteredGames
}

func FormatGameTitle(home Team, away Team, withRank bool) string {
	if withRank {
		switch {
		case home.Rank < 26 && away.Rank < 26:
			// both teams ranked
			return fmt.Sprintf("%s %s (%d) v %s %s (%d)\n",
				home.Location, home.Name, home.Rank, away.Location, away.Name, away.Rank)
		case home.Rank > 25 && away.Rank < 26:
			// home team not ranked, away team ranked
			return fmt.Sprintf("%s %s v %s %s (%d)\n",
				home.Location, home.Name, away.Location, away.Name, away.Rank)
		case home.Rank < 26 && away.Rank > 25:
			// home team ranked, away team not ranked
			return fmt.Sprintf("%s %s (%d) v %s %s\n",
				home.Location, home.Name, home.Rank, away.Location, away.Name)
		}
	}
	return fmt.Sprintf("%s %s v %s %s\n",
		home.Location, home.Name, away.Location, away.Name)
}

func FormatTeamDetail(team Team) string {
	if team.FullRecord == "" {
		team.FullRecord = "N/A"
	}
	if team.HomeRecord == "" {
		team.HomeRecord = "N/A"
	}
	if team.AwayRecord == "" {
		team.AwayRecord = "N/A"
	}
	if team.ConferenceRecord == "" {
		team.ConferenceRecord = "N/A"
	}
	return fmt.Sprintf("%s %s Season Record: %s, Home: %s, Away: %s, Conference: %s\n",
		team.Location, team.Name, team.FullRecord, team.HomeRecord, team.AwayRecord, team.ConferenceRecord)
}

func FormatStartTime(start string) string {
	return fmt.Sprintf("Start time: %s CST\n", start)
}

func FormatStatus(status string) string {
	return fmt.Sprintf("Status: %s\n", status)
}

func FormatGameScore(home Team, away Team) string {
	return fmt.Sprintf("Scores: %s - %s\n", home.Score, away.Score)
}
