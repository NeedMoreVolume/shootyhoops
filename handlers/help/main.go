package help

const helpMessage = `Halp?
General
	by even reaching this menu, I assume you know you must mention me to get me to respond.
	You can change my avatar now by using the following command
	'set-avatar <url>' - I will unhappily respond with a series of questions and something Rick Adelman once said as a coach if you do not provide a valid url for this
NCAA
	'ncaa-games' - will return mens NCAA games for today
	'ncaa-games tmrw' - will return mens NCAA games for tomorrow
	'ncaa-games tomorrow' - will also return NCAA games for tomorrow
	'ncaa-games YYYYMMDD' - will return mens NCAA games for the given date
NBA
	'nba-games' - will return NBA games for today
	'nba-games tmrw' - will return NBA games for tomorrow
	'nba-games tomorrow' - will also return NBA games for tomorrow
	'nba-games YYYYMMDD' - will return NBA games for the given date
Extra
	I can provide some extra detail with special flags, but please put the flags after the commands above.
	Current flags I support
		-scores - I will give you the current score of the games listed, if I can.
		-detailed - I will give you extra record info about the teams playing (total, home, away and conference records)
		-team - get data for a specified team
Coming Soon...
	'ncaa-standings' - will return ncaa standings
	'ncaa-standings <conference>' - will return ncaa standings of a given conference
	'nba-standings' - will return nba standings
	'nba-standings east' - returns Eastern Conference standings
	'nba-standings west' - returns Western Conference standings
`

func Help() string {
	return helpMessage
}
