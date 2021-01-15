package espn

type Type struct {
	ID           string `json:"id"`
	Type         int    `json:"type"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type Season struct {
	Year      int    `json:"year"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type League struct {
	ID                  string   `json:"id"`
	UID                 string   `json:"uid"`
	Name                string   `json:"name"`
	Abbreviation        string   `json:"abbreviation"`
	MidsizeName         string   `json:"midsizeName"`
	Slug                string   `json:"slug"`
	Season              Season   `json:"season"`
	CalendarType        string   `json:"calendarType"`
	CalendarIsWhitelist bool     `json:"calendarIsWhitelist"`
	CalendarStartDate   string   `json:"calendarStartDate"`
	CalendarEndDate     string   `json:"calendarEndDate"`
	Calendar            []string `json:"calendar"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Venue struct {
	ID       string  `json:"id"`
	FullName string  `json:"fullName"`
	Address  Address `json:"address"`
	Capacity int     `json:"capacity"`
	Indoor   bool    `json:"indoor"`
}

type Competitor struct {
	ID       string `json:"id"`
	UID      string `json:"uid"`
	Type     string `json:"type"`
	Order    int    `json:"order"`
	HomeAway string `json:"homeAway"`
	Winner   bool   `json:"winner"`
	Team     struct {
		ID               string `json:"id"`
		UID              string `json:"uid"`
		Location         string `json:"location"`
		Name             string `json:"name"`
		Abbreviation     string `json:"abbreviation"`
		DisplayName      string `json:"displayName"`
		ShortDisplayName string `json:"shortDisplayName"`
		Color            string `json:"color"`
		AlternateColor   string `json:"alternateColor"`
		IsActive         bool   `json:"isActive"`
		Venue            struct {
			ID string `json:"id"`
		} `json:"venue"`
		Links []struct {
			Rel        []string `json:"rel"`
			Href       string   `json:"href"`
			Text       string   `json:"text"`
			IsExternal bool     `json:"isExternal"`
			IsPremium  bool     `json:"isPremium"`
		} `json:"links"`
		Logo         string `json:"logo"`
		ConferenceID string `json:"conferenceId"`
	} `json:"team"`
	Score      string `json:"score"`
	Linescores []struct {
		Value float64 `json:"value"`
	} `json:"linescores"`
	Statistics []struct {
		Name         string `json:"name"`
		Abbreviation string `json:"abbreviation"`
		DisplayValue string `json:"displayValue"`
	} `json:"statistics"`
	CuratedRank struct {
		Current int `json:"current"`
	} `json:"curatedRank"`
	Records []struct {
		Name         string `json:"name"`
		Abbreviation string `json:"abbreviation,omitempty"`
		Type         string `json:"type"`
		Summary      string `json:"summary"`
	} `json:"records"`
}

type Competition struct {
	ID                    string        `json:"id"`
	UID                   string        `json:"uid"`
	Date                  string        `json:"date"`
	Attendance            int           `json:"attendance"`
	Type                  Type          `json:"type"`
	TimeValid             bool          `json:"timeValid"`
	NeutralSite           bool          `json:"neutralSite"`
	ConferenceCompetition bool          `json:"conferenceCompetition"`
	Recent                bool          `json:"recent"`
	Venue                 Venue         `json:"venue"`
	Competitors           []Competitor  `json:"competitors"`
	Notes                 []interface{} `json:"notes"`
	Status                struct {
		Clock        float64 `json:"clock"`
		DisplayClock string  `json:"displayClock"`
		Period       int     `json:"period"`
		Type         struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			State       string `json:"state"`
			Completed   bool   `json:"completed"`
			Description string `json:"description"`
			Detail      string `json:"detail"`
			ShortDetail string `json:"shortDetail"`
		} `json:"type"`
	} `json:"status"`
	Broadcasts []struct {
		Market string   `json:"market"`
		Names  []string `json:"names"`
	} `json:"broadcasts"`
	StartDate     string `json:"startDate"`
	GeoBroadcasts []struct {
		Type struct {
			ID        string `json:"id"`
			ShortName string `json:"shortName"`
		} `json:"type"`
		Market struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"market"`
		Media struct {
			ShortName string `json:"shortName"`
		} `json:"media"`
		Lang   string `json:"lang"`
		Region string `json:"region"`
	} `json:"geoBroadcasts"`
}

type Event struct {
	ID           string        `json:"id"`
	UID          string        `json:"uid"`
	Date         string        `json:"date"`
	Name         string        `json:"name"`
	ShortName    string        `json:"shortName"`
	Season       Season        `json:"season"`
	Competitions []Competition `json:"competitions"`
	Links        []struct {
		Language   string   `json:"language"`
		Rel        []string `json:"rel"`
		Href       string   `json:"href"`
		Text       string   `json:"text"`
		ShortText  string   `json:"shortText"`
		IsExternal bool     `json:"isExternal"`
		IsPremium  bool     `json:"isPremium"`
	} `json:"links"`
	Status struct {
		Clock        float64 `json:"clock"`
		DisplayClock string  `json:"displayClock"`
		Period       int     `json:"period"`
		Type         struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			State       string `json:"state"`
			Completed   bool   `json:"completed"`
			Description string `json:"description"`
			Detail      string `json:"detail"`
			ShortDetail string `json:"shortDetail"`
		} `json:"type"`
	} `json:"status"`
}

type EventsDate struct {
	Date       string `json:"date"`
	SeasonType int    `json:"seasonType"`
}

type GameResponse struct {
	Leagues    []League   `json:"leagues"`
	Events     []Event    `json:"events"`
	EventsDate EventsDate `json:"eventsDate"`
}

type StandingResponse struct {
	UID          string     `json:"uid"`
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Abbreviation string     `json:"abbreviation"`
	Children     []Children `json:"children"`
	Links        []Links    `json:"links"`
	Seasons      []Seasons  `json:"seasons"`
}

type Links struct {
	Language   string   `json:"language"`
	Rel        []string `json:"rel"`
	Href       string   `json:"href"`
	Text       string   `json:"text"`
	ShortText  string   `json:"shortText"`
	IsExternal bool     `json:"isExternal"`
	IsPremium  bool     `json:"isPremium"`
}

type Logos struct {
	Href   string   `json:"href"`
	Width  int      `json:"width"`
	Height int      `json:"height"`
	Alt    string   `json:"alt"`
	Rel    []string `json:"rel"`
}

type Team struct {
	ID               string  `json:"id"`
	UID              string  `json:"uid"`
	Location         string  `json:"location"`
	Name             string  `json:"name"`
	Abbreviation     string  `json:"abbreviation"`
	DisplayName      string  `json:"displayName"`
	ShortDisplayName string  `json:"shortDisplayName"`
	IsActive         bool    `json:"isActive"`
	Logos            []Logos `json:"logos"`
	Links            []Links `json:"links"`
}

type Stats struct {
	Name             string  `json:"name"`
	DisplayName      string  `json:"displayName"`
	ShortDisplayName string  `json:"shortDisplayName"`
	Description      string  `json:"description"`
	Abbreviation     string  `json:"abbreviation,omitempty"`
	Type             string  `json:"type"`
	Value            float64 `json:"value,omitempty"`
	DisplayValue     string  `json:"displayValue"`
	ID               string  `json:"id,omitempty"`
	Summary          string  `json:"summary,omitempty"`
}

type Entries struct {
	Team  Team    `json:"team"`
	Stats []Stats `json:"stats"`
}

type Standings struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"displayName"`
	Links       []Links   `json:"links"`
	Season      int       `json:"season"`
	SeasonType  int       `json:"seasonType"`
	Entries     []Entries `json:"entries"`
}

type Children struct {
	UID          string    `json:"uid"`
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Abbreviation string    `json:"abbreviation"`
	Standings    Standings `json:"standings"`
}

type Types struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	HasStandings bool   `json:"hasStandings"`
}

type Seasons struct {
	Year        int     `json:"year"`
	StartDate   string  `json:"startDate"`
	EndDate     string  `json:"endDate"`
	DisplayName string  `json:"displayName"`
	Types       []Types `json:"types"`
}
