package f1

import "context"

const (
	teamsEndpoint = "/players"
)

// TeamsService handles communication with the teams related
// methods of the API.
type TeamsService service

type Team struct {
	ID                          int                    `json:"id"`
	Name                        string                 `json:"first_name"`
	PositionAbbreviation        string                 `json:"position_abbreviation"`
	Price                       float64                `json:"price"`
	CurrentPriceChangeInfo      CurrentPriceChangeInfo `json:"current_price_change_info"`
	Banned                      bool                   `json:"banned"`
	StreakEventsProgress        StreakEventsProgress   `json:"streak_events_progress"`
	TeamAbbreviation            string                 `json:"team_abbreviation"`
	WeeklyPriceChange           float64                `json:"weekly_price_change"`
	WeeklyPriceChangePercentage int                    `json:"weekly_price_change_percentage"`
	TeamID                      int                    `json:"team_id"`
	Headshot                    Headshot               `json:"headshot"`
	Score                       int                    `json:"score"`
	IsConstructor               bool                   `json:"is_constructor"`
	SeasonScore                 float64                `json:"season_score"`
	ConstructorData             ConstructorData        `json:"constructor_data"`
	SeasonPrices                []Price                `json:"season_prices"`
	NumFixturesInGameweek       int                    `json:"num_fixtures_in_gameweek"`
	DeletedInFeed               bool                   `json:"deleted_in_feed"`
	HasFixture                  bool                   `json:"has_fixture"`
	DisplayName                 string                 `json:"display_name"`
	ExternalID                  string                 `json:"external_id"`
	ProfileImage                ProfileImage           `json:"profile_image"`
}

type ConstructorData struct {
	BestFinish          int     `json:"best_finish"`
	BestFinishCount     int     `json:"best_finish_count"`
	BestGrid            int     `json:"best_grid"`
	BestGridCount       int     `json:"best_grid_count"`
	Titles              int     `json:"titles"`
	ChampionshipPoints  float64 `json:"championship_points"`
	FirstSeason         string  `json:"first_season"`
	Poles               int     `json:"poles"`
	FastestLaps         int     `json:"fastest_laps"`
	Country             string  `json:"country"`
	HighestRaceFinished string  `json:"highest_race_finished"`
}

// ListAll lists all teams.
func (s *TeamsService) ListAll(ctx context.Context) ([]*Team, error) {
	url := s.client.baseURL.String() + teamsEndpoint
	var res struct {
		Data []*Team `json:"players"`
	}
	err := s.client.get(ctx, url, &res)
	if err != nil {
		return nil, err
	}

	teams := make([]*Team, 0, 10)

	for _, team := range res.Data {
		if team.IsConstructor {
			teams = append(teams, team)
		}
	}

	return teams, nil
}

// Get a single team by id.
func (s *TeamsService) GetOne(ctx context.Context, id int) (*Team, error) {
	url := s.client.baseURL.String() + teamsEndpoint
	var res struct {
		Data []*Team `json:"players"`
	}
	err := s.client.get(ctx, url, &res)
	if err != nil {
		return nil, err
	}

	var team *Team
	for _, t := range res.Data {
		if t.ID == id {
			if !t.IsConstructor {
				break
			}
			team = t
			break
		}
	}

	return team, nil
}
