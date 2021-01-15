package nhl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shootyhoops/models"
	"shootyhoops/models/espn"
	"strings"
	"time"
)

const nhlEspnScoreboardUrl = "https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/scoreboard"

//const nhlEspnStandingsUrl = "https://site.api.espn.com/apis/v2/sports/basketball/mens-college-basketball/standings"

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

	var teamName *string
	if strings.Contains(message, "-team") {
		teamSplit := strings.Split(message, "-team")
		split := strings.Split(teamSplit[1], " ")
		message = strings.Replace(message, "-team", "", 1)

		// assign team name and remove from main message
		tn := split[1]
		message = strings.Replace(message, tn, "", 1)
		teamName = &tn
	}

	message = strings.TrimSpace(message)
	dateCheck := strings.Split(message, " ")
	query := time.Now().Format("20060102")

	if len(dateCheck) > 1 {
		// today is handled by default in the setting of the query variable.
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

	url := nhlEspnScoreboardUrl + "?dates=" + query
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

	return models.GamesToMessage(models.EspnGamesToGames(gamesData), true, detailed, scores, teamName)
}
