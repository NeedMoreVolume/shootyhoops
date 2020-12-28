package nba

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shootyhoops/models"
	"shootyhoops/models/espn"
	"strings"
	"time"

	"github.com/NeedMoreVolume/shootyhoops/models"
	"github.com/NeedMoreVolume/shootyhoops/models/espn"
)

const nbaEspnScoreboardUrl = "https://site.api.espn.com/apis/site/v2/sports/basketball/nba/scoreboard"
const nbaEspnStandingsUrl = "https://site.api.espn.com/apis/v2/sports/basketball/nba/standings"

func Games(message string) string {
	var detailed bool
	if strings.Contains(message, "-detailed") {
		message = strings.Replace(message, "-detailed", "", 1)
		detailed = true
	}

	var scores bool
	if strings.Contains(message, "-scores") {
		message = strings.Replace(message, "-scores", "", 1)
		scores = true
	}

	var teamName string
	if strings.Contains(message, "-team") {
		split := strings.Split(message, " ")
		message = strings.Replace(message, "-team", "", 1)

		teamIndex := 0
		for i := range split {
			if split[i] != "-team" {
				continue
			}

			teamIndex = i
			break
		}

		// assign team name and remove from main message
		teamName = split[teamIndex+1]
		message = strings.Replace(message, teamName, "", 1)
	}

	message = strings.TrimSpace(message)
	dateCheck := strings.Split(message, " ")
	query := time.Now().Format("20060102")

	if len(dateCheck) > 1 {
		switch dateCheck[1] {
		case "today":
		case "tmrw", "tomorrow":
			query = time.Now().Add(time.Hour * 24).Format("20060102")
		default:
			// trust... but verify! :D
			queryTime, err := time.Parse("20060102", dateCheck[1])
			if err != nil {
				return "invalid date, the required format is YYYYMMDD."
			}
			query = queryTime.Format("20060102")
		}
	}

	url := nbaEspnScoreboardUrl + "?dates=" + query
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return "@NeedMoreVolume, please fix me. I've fallen and I can't get the ncaa game data."
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "@NeedMoreVolume, what the fuck man this isn't a readable body, who the fuck did you tell me to talk to for game data?"
	}

	var gamesData espn.GameResponse
	err = json.Unmarshal(body, &gamesData)
	if err != nil {
		fmt.Println(err.Error())
		return "@NeedMoreVolume, either you are a dumbass or the espn response changed."
	}

	// filter for specific team if requested
	if teamName != "" {
		gamesData.Events = FilterTeams(gamesData.Events, teamName)
	}
	return models.GamesToMessage(models.EspnGamesToGames(gamesData), false, detailed, scores)
}

func FilterTeams(events []espn.Event, teamName string) []espn.Event {
	filteredEvents := []espn.Event{}
	for i := range events {
		var breakout bool
		competitors := events[i].Competitions[0].Competitors
		for j := range competitors {
			// todo: this needs to be more concise to avoid false positives
			// i.e. 'nets' also returning 'hornets'
			// use a regex identifier to do it probs
			espnName := strings.ToLower(competitors[j].Team.DisplayName)
			if !strings.Contains(espnName, strings.ToLower(teamName)) {
				continue
			}

			filteredEvents = append(filteredEvents, events[i])
			breakout = true
			break
		}

		if breakout {
			break
		}
	}

	return filteredEvents
}

func Standings(message string) string {
	url := nbaEspnStandingsUrl
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return "@NeedMoreVolume, please fix me. I've fallen and I can't get the ncaa game data."
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "@NeedMoreVolume, what the fuck man this isn't a readable body, who the fuck did you tell me to talk to for game data?"
	}

	var standingsData espn.StandingResponse
	err = json.Unmarshal(body, &standingsData)
	if err != nil {
		fmt.Println(err.Error())
		return "@NeedMoreVolume, either you are a dumbass or the espn response changed."
	}

	messageSplit := strings.Split(message, " ")
	if len(messageSplit) > 1 {
		switch messageSplit[1] {
		case "east", "eastern":
			// only get eastern conference standings data
		case "west", "western":
			// only get western conference standings data
		default:
			// give them all the conference standings data
		}
	}

	return "implementation not complete"
}
