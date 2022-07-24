package f1

import (
	"context"
	"time"
)

const (
	circuitsEndpoint = "/circuits"
)

// CircuitsService handles communication with the Circuits related
// methods of the API.
type CircuitsService service

type Circuit struct {
	GamePeriodID   int    `json:"game_period_id"`
	GamePeriodName string `json:"game_period_name"`
	Circuit        struct {
		ID             int       `json:"id"`
		FirstGrandPrix string    `json:"first_grand_prix"`
		LapsTotal      string    `json:"laps_total"`
		Length         string    `json:"length"`
		Distance       string    `json:"distance"`
		LapRecord      string    `json:"lap_record"`
		PraticeOne     string    `json:"pratice_one"`
		PraticeTwo     string    `json:"pratice_two"`
		PraticeThree   string    `json:"pratice_three"`
		Qualifying     string    `json:"qualifying"`
		Race           string    `json:"race"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		Name           string    `json:"name"`
		ShortName      string    `json:"short_name"`
		CountryIso     string    `json:"country_iso"`
		CircuitImage   struct {
			URL string `json:"url"`
		} `json:"circuit_image"`
		SprintQualifying *string `json:"sprint_qualifying"`
		GmtOffset        string  `json:"gmt_offset"`
	} `json:"circuit"`
}

// ListAll lists all circuits.
func (s *CircuitsService) ListAll(ctx context.Context) ([]*Circuit, error) {
	url := s.client.baseURL.String() + circuitsEndpoint
	var res []*Circuit
	err := s.client.get(ctx, url, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Get a single circuit by id.
func (s *CircuitsService) GetOne(ctx context.Context, id int) (*Circuit, error) {
	circuits, err := s.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	var circuit *Circuit
	for _, c := range circuits {
		if c.Circuit.ID == id {
			circuit = c
			break
		}
	}

	return circuit, nil
}
