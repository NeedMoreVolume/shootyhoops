package help

const helpMessage = "Halp?\n``` General\nYou can change my avatar now by using the following command\n	'set-avatar <url>' - I will unhappily respond with a series of questions and something Rick Adelman once said as a coach if you do not provide a valid url for this```\n``` NCAA\n	'ncaa-games' - will return mens NCAA games for today\n	'ncaa-games tmrw' - will return mens NCAA games for tomorrow\n	'ncaa-games tomorrow' - will also return NCAA games for tomorrow\n	'ncaa-games YYYYMMDD' - will return mens NCAA games for the given date```\n``` NBA\n	'nba-games' - will return NBA games for today\n	'nba-games tmrw' - will return NBA games for tomorrow\n	'nba-games tomorrow' - will also return NBA games for tomorrow\n	'nba-games YYYYMMDD' - will return NBA games for the given date```\n``` Extra\nI can provide some extra detail with special flags, but please put the flags after the commands above.\nCurrent flags I support\n	-scores - I will give you the current score of the games listed, if I can.\n	-detailed - I will give you extra record info about the teams playing (total, home, away and conference records)\n	-team - get data for a specified team```\n```Coming Soon... maybe\n	'ncaa-standings' - will return ncaa standings\n	'ncaa-standings <conference>' - will return ncaa standings of a given conference\n	'nba-standings' - will return nba standings\n	'nba-standings east' - returns Eastern Conference standings\n	'nba-standings west' - returns Western Conference standings```\n"

func Help() string {
	return helpMessage
}
