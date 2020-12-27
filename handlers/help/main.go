package help

const helpMessage = `Halp?
General
	by even reaching this menu, I assume you know you must mention me to get me to respond.
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
Coming Soon...
	ncaa-standings - will return ncaa standings
	nba-standings - will return nba standings
	nba-standings east - returns Eastern Conference standings
	nba-standings west - returns Western Conference standings
`

func Help() string {
	return helpMessage
}
