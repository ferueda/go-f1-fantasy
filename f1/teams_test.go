package f1

import (
	"context"
	"testing"
)

const (
	numOfTeams = 10
)

func TestTeamsService_ListAll(t *testing.T) {
	c := NewClient(nil)
	teams, err := c.Teams.ListAll(context.Background())
	if err != nil {
		t.Errorf("Teams.ListAll returned error: %v", err)
	}

	if len(teams) != numOfTeams {
		t.Errorf("Teams.ListAll returned %v teams, want %v", len(teams), numOfTeams)
	}

	for _, team := range teams {
		if !team.IsConstructor {
			t.Errorf("Teams.ListAll returned driver %v with ID %v. No driver should be returned.", team.DisplayName, team.ID)
		}
	}
}

func TestTeamsService_GetOne(t *testing.T) {
	c := NewClient(nil)
	t1Id := 2
	t1, err := c.Teams.GetOne(context.Background(), t1Id)
	if err != nil {
		t.Errorf("Teams.GetOne returned error: %v", err)
	}

	if t1 == nil {
		t.Errorf("Teams.GetOne returned nil, want team with id %v", t1Id)
	}

	if t1.ID != t1Id {
		t.Errorf("Teams.GetOne returned team with id %v, want team with id %v", t1.ID, t1Id)
	}

	if !t1.IsConstructor {
		t.Errorf("Teams.GetOne returned driver %v with ID %v. want team with id %v.", t1.DisplayName, t1.ID, t1Id)
	}

	t2, err := c.Teams.GetOne(context.Background(), 12)
	if err != nil {
		t.Errorf("Teams.GetOne returned error: %v", err)
	}

	if t2 != nil {
		t.Errorf("Teams.GetOne returned driver/team %v with ID %v, want nil", t2.DisplayName, t2.ID)
	}

	dt3, err := c.Teams.GetOne(context.Background(), 1000)
	if err != nil {
		t.Errorf("Teams.GetOne returned error: %v", err)
	}

	if dt3 != nil {
		t.Errorf("Teams.GetOne returned driver/team %v with ID %v, want nil", dt3.DisplayName, dt3.ID)
	}
}
