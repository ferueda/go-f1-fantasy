package f1

import "context"

const (
	driversEndpoint = "/players"
)

// DriversService handles communication with the Drivers related
// methods of the API.
type DriversService service

type Driver struct {
	ID                          int                    `json:"id"`
	FirstName                   string                 `json:"first_name"`
	LastName                    string                 `json:"last_name"`
	TeamName                    string                 `json:"team_name"`
	Price                       float64                `json:"price"`
	CurrentPriceChangeInfo      CurrentPriceChangeInfo `json:"current_price_change_info"`
	Injured                     bool                   `json:"injured"`
	Banned                      bool                   `json:"banned"`
	StreakEventsProgress        StreakEventsProgress   `json:"streak_events_progress"`
	ChanceOfPlaying             float64                `json:"chance_of_playing"`
	TeamAbbreviation            string                 `json:"team_abbreviation"`
	WeeklyPriceChange           float64                `json:"weekly_price_change"`
	WeeklyPriceChangePercentage int                    `json:"weekly_price_change_percentage"`
	TeamID                      int                    `json:"team_id"`
	Headshot                    Headshot               `json:"headshot"`
	Score                       int                    `json:"score"`
	ShirtNumber                 int                    `json:"shirt_number"`
	Country                     string                 `json:"country"`
	CountryIso                  string                 `json:"country_iso"`
	IsConstructor               bool                   `json:"is_constructor"`
	SeasonScore                 float64                `json:"season_score"`
	DriverData                  DriverData             `json:"driver_data"`
	BornAt                      string                 `json:"born_at"`
	SeasonPrices                []Price                `json:"season_prices"`
	NumFixturesInGameweek       int                    `json:"num_fixtures_in_gameweek"`
	DeletedInFeed               bool                   `json:"deleted_in_feed"`
	HasFixture                  bool                   `json:"has_fixture"`
	DisplayName                 string                 `json:"display_name"`
	ExternalID                  string                 `json:"external_id"`
	ProfileImage                ProfileImage           `json:"profile_image"`
	MiscImage                   MiscImage              `json:"misc_image"`
}

type CurrentPriceChangeInfo struct {
	CurrentSelectionPercentage     float64 `json:"current_selection_percentage"`
	ProbabilityPriceUpPercentage   float64 `json:"probability_price_up_percentage"`
	ProbabilityPriceDownPercentage float64 `json:"probability_price_down_percentage"`
}

type StreakEventsProgress struct {
	TopTenInARowQualifyingProgress int `json:"top_ten_in_a_row_qualifying_progress"`
	TopTenInARowRaceProgress       int `json:"top_ten_in_a_row_race_progress"`
}

type Headshot struct {
	Profile    string `json:"profile"`
	PitchView  string `json:"pitch_view"`
	PlayerList string `json:"player_list"`
}

type DriverData struct {
	Wins                int    `json:"wins"`
	Podiums             int    `json:"podiums"`
	Poles               int    `json:"poles"`
	FastestLaps         int    `json:"fastest_laps"`
	GrandsPrixEntered   int    `json:"grands_prix_entered"`
	Titles              int    `json:"titles"`
	ChampionshipPoints  int    `json:"championship_points"`
	BestFinish          int    `json:"best_finish"`
	BestFinishCount     int    `json:"best_finish_count"`
	BestGrid            int    `json:"best_grid"`
	BestGridCount       int    `json:"best_grid_count"`
	HighestRaceFinished string `json:"highest_race_finished"`
	PlaceOfBirth        string `json:"place_of_birth"`
}

type Price struct {
	GamePeriodID int     `json:"game_period_id"`
	Price        float64 `json:"price"`
}

type ProfileImage struct {
	URL string `json:"url"`
}

type MiscImage struct {
	URL interface{} `json:"url"`
}

// ListAll lists all drivers.
func (s *DriversService) ListAll(ctx context.Context) ([]*Driver, error) {
	url := s.client.baseURL.String() + driversEndpoint
	var res struct {
		Data []*Driver `json:"players"`
	}
	err := s.client.get(ctx, url, &res)
	if err != nil {
		return nil, err
	}

	drivers := make([]*Driver, 0, 20)

	for _, driver := range res.Data {
		if !driver.IsConstructor {
			drivers = append(drivers, driver)
		}
	}

	return drivers, nil
}

// Get a single driver by id.
func (s *DriversService) GetOne(ctx context.Context, id int) (*Driver, error) {
	url := s.client.baseURL.String() + driversEndpoint
	var res struct {
		Data []*Driver `json:"players"`
	}
	err := s.client.get(ctx, url, &res)
	if err != nil {
		return nil, err
	}

	var driver *Driver
	for _, d := range res.Data {
		if d.ID == id {
			if d.IsConstructor {
				break
			}
			driver = d
			break
		}
	}

	return driver, nil
}
